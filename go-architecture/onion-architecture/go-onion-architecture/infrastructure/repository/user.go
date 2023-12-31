package infrastructure

import (
	"time"

	model "go-onion-architecture-sample/domain/model"
	repository "go-onion-architecture-sample/domain/repository"
	db "go-onion-architecture-sample/infrastructure"
)

type userRepository struct {
	dbClient *db.DBClient
}

func NewUserRepository(db *db.DBClient) repository.UserRepository {
	return &userRepository{
		dbClient: db,
	}
}

func (u *userRepository) Create(name string) (model.UserId, error) {
	stmt, err := u.dbClient.Client.Prepare("INSERT INTO users (name, created_at) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return model.UserId(id), nil
}

func (u *userRepository) Get(id model.UserId) (model.User, error) {
	stmt, err := u.dbClient.Client.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.CreatedAt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
