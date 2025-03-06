package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"golang_startup/internal/model"
)

// FindUser - Returns a user for a given ID or an error if anything fails
func (db *Database) FindUser(ctx context.Context, id string) (model.User, error) {
	row := db.client.QueryRowContext(ctx, `
		SELECT id, username, email, created, updated
		FROM users WHERE id = $1
		LIMIT 1
	`, id)

	var (
		uid, name, email sql.NullString
		created, updated sql.NullTime
	)

	err := row.Scan(&uid, &name, &email, &created, &updated)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		ID:        uid.String,
		Name:      name.String,
		Email:     email.String,
		CreatedAt: created.Time,
		UpdatedAt: updated.Time,
	}

	return user, nil
}

// UpdateUser - Updates a user and returns an error if anything fails
func (db *Database) UpdateUser(ctx context.Context, user model.User) error {
	err := db.client.QueryRowContext(ctx, `
		UPDATE users 
		SET username = $2, email = $3
		WHERE id = $1
	`, user.ID, user.Name, user.Email).Err()
	if err != nil {
		return err
	}

	return nil
}

// CreateUser - Creates a user and returns its ID or an error, if anything fails
func (db *Database) CreateUser(ctx context.Context, user model.User) (string, error) {
	lastInsertedID := ""

	err := db.client.QueryRowContext(ctx, `
		INSERT INTO users (id, username, email)
		VALUES ($1, $2, $3)
		RETURNING id
	`, user.ID, user.Name, user.Email).Scan(&lastInsertedID)
	if err != nil {
		return "", err
	}

	return lastInsertedID, nil
}

// DeleteUser - Deletes a User for a given ID and could return an error if anything fails
func (db *Database) DeleteUser(ctx context.Context, id string) error {
	err := db.client.QueryRowContext(ctx, `
		DELETE FROM users
		WHERE id = $1
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
		LIMIT $1
	`, limit)
	if err != nil {
		return []model.User{}, fmt.Errorf("error executing query: %s", err.Error())
	}

	if rows.Err() != nil {
		return []model.User{}, fmt.Errorf("rows returned an error: %s", rows.Err())
	}

	var (
		uid, name, email sql.NullString
		created, updated sql.NullTime
		users            []model.User
	)

	for rows.Next() {
		err = rows.Scan(&uid, &name, &email, &created, &updated)
		if err != nil {
			return []model.User{}, fmt.Errorf("error scanning rows: %s", err.Error())
		}

		user := model.User{
			ID:        uid.String,
			Name:      name.String,
			Email:     email.String,
			CreatedAt: created.Time,
			UpdatedAt: updated.Time,
		}

		users = append(users, user)
	}

	return users, nil
}
