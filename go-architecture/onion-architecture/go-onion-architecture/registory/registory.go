//go:generate mockgen -source=$GOFILE -destination=../mocks/$GOPACKAGE/mock_$GOFILE -package=mock_registory

package registory

import (
	service "go-onion-architecture-sample/application/service"
	repository "go-onion-architecture-sample/domain/repository"
	db "go-onion-architecture-sample/infrastructure"
	infrastructure "go-onion-architecture-sample/infrastructure/repository"
)

type Registory interface {
	UserRep() repository.UserRepository
	UserService() service.UserService
	TodoRep() repository.TodoRepository
	TodoService() service.TodoService
}

type registory struct {
	dbClient *db.DBClient
}

func NewRegistory(db *db.DBClient) Registory {
	return &registory{
		dbClient: db,
	}
}

func (r *registory) UserRep() repository.UserRepository {
	return infrastructure.NewUserRepository(r.dbClient)
}

func (r *registory) UserService() service.UserService {
	return service.NewUserService(r.UserRep())
}

func (r *registory) TodoRep() repository.TodoRepository {
	return infrastructure.NewTodoRepository(r.dbClient)
}

func (r *registory) TodoService() service.TodoService {
	return service.NewTodoService(r.TodoRep())
}
