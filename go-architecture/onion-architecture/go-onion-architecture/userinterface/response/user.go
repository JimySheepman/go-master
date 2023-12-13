package response

import "go-onion-architecture-sample/domain/model"

type CreateUserResponse struct {
	Id model.UserId `json:"id"`
}

type GetUserResponse struct {
	User model.User `json:"user"`
}
