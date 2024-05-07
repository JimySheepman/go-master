//go:generate mockgen -source=$GOFILE -destination=../../mocks/mock_$GOPACKAGE/$GOFILE

package repository

import "go-onion-architecture-sample/domain/model"

type TodoRepository interface {
	Create(model.UserId, model.Title, model.Description) (model.TodoId, error)
	List(model.UserId) ([]model.Todo, error)
	Delete(model.TodoId) error
}
