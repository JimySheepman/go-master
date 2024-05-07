package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"golang_startup/internal/model"
)

const (
	baseNumber = 10

	timeLayout = "2006-01-02 15:04:05"
)

// FindUser - Returns a user for a given ID or an error if anything fails
func (db *Database) FindUser(ctx context.Context, id string) (model.User, error) {
	row := db.client.QueryRowContext(ctx, `
		SELECT id, username, email, created, updated
		FROM users WHERE id = ?
		LIMIT 1
	`, id)

	var uid, name, email, created, updated sql.NullString

	err := row.Scan(&uid, &name, &email, &created, &updated)
	if err != nil {
		return model.User{}, err
	}

	createdTime, err := time.Parse(timeLayout, created.String)
	if err != nil {
		return model.User{}, err
	}

	updatedTime, err := time.Parse(timeLayout, updated.String)
	if err != nil {
		return model.User{}, fmt.Errorf("error parsing time layout: %s", err.Error())
	}

	user := model.User{
		ID:        uid.String,
		Name:      name.String,
		Email:     email.String,
		CreatedAt: createdTime,
		UpdatedAt: updatedTime,
	}

	return user, nil
}

// UpdateUser - Updates a user and returns an error if anything fails
func (db *Database) UpdateUser(ctx context.Context, user model.User) error {
	err := db.client.QueryRowContext(ctx, `
		UPDATE users 
		SET username = ?, email = ?
		WHERE id = ?
	`, user.Name, user.Email, user.ID).Err()
	if err != nil {
		return err
	}

	return nil
}

// CreateUser - Creates a user and returns its ID or an error, if anything fails
func (db *Database) CreateUser(ctx context.Context, user model.User) (string, error) {
	res, err := db.client.ExecContext(ctx, `
		INSERT INTO users (username, email)
		VALUES (?, ?)
	`, user.Name, user.Email)
	if err != nil {
		return "", fmt.Errorf("error creating user: %s", err.Error())
	}

	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("error returning last insert ID: %s", err.Error())
	}

	log.Println(lastInsertedID)

	return strconv.FormatInt(lastInsertedID, baseNumber), nil
}

// DeleteUser - Deletes a User for a given ID and could return an error if anything fails
func (db *Database) DeleteUser(ctx context.Context, id string) error {
	err := db.client.QueryRowContext(ctx, `
		DELETE FROM users
		WHERE id = ?
	`, id).Err()
	if err != nil {
		return err
	}

	return nil
}

// FindAllUsers - Returns all users for a given limit and offset
func (db *Database) FindAllUsers(ctx context.Context, offset int64, limit int64) ([]model.User, error) {
	rows, err := db.client.QueryContext(ctx, `
		SELECT id, username, email, created, updated
		FROM users
		LIMIT ?
	`, limit)
	if err != nil {
		return []model.User{}, fmt.Errorf("error executing query: %s", err.Error())
	}

	if rows.Err() != nil {
		return []model.User{}, fmt.Errorf("rows returned an error: %s", rows.Err())
	}

	var (
		uid, name, email, created, updated sql.NullString
		users                              []model.User
	)

	for rows.Next() {
		err = rows.Scan(&uid, &name, &email, &created, &updated)
		if err != nil {
			return []model.User{}, fmt.Errorf("error scanning rows: %s", err.Error())
		}

		createdTime, err := time.Parse(timeLayout, created.String)
		if err != nil {
			return []model.User{}, fmt.Errorf("error parsing time layout: %s", err.Error())
		}

		updatedTime, err := time.Parse(timeLayout, updated.String)
		if err != nil {
			return []model.User{}, fmt.Errorf("error parsing time layout: %s", err.Error())
		}

		user := model.User{
			ID:        uid.String,
			Name:      name.String,
			Email:     email.String,
			CreatedAt: createdTime,
			UpdatedAt: updatedTime,
		}

		users = append(users, user)
	}

	return users, nil
}
