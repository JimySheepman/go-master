package response

import "go-onion-architecture-sample/domain/model"

type CreateTodoResponse struct {
	Id model.TodoId `json:"id"`
}

type GetTodosResponse struct {
	Todos []model.Todo `json:"todos"`
}
