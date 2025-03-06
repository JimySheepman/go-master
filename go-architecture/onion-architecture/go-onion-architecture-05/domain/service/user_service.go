package service

import (
	"context"

	"go-echo-rest-sample/domain/repository"
)

type UserService interface {
	DoSomething(ctx context.Context, foo int) error
}

type userService struct {
	repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (u *userService) DoSomething(ctx context.Context, foo int) error {
	// some code
	return nil
}
