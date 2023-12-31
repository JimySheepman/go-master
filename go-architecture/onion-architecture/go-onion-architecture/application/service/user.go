//go:generate mockgen -source=$GOFILE -destination=../../mocks/service/mock_$GOFILE -package=mock_service

package service

import (
	"go-onion-architecture-sample/domain/model"
	"go-onion-architecture-sample/domain/repository"
)

type UserService interface {
	Create(string) (model.UserId, error)
	Get(model.UserId) (model.User, error)
}

type userService struct {
	userRep repository.UserRepository
}

func NewUserService(userRep repository.UserRepository) UserService {
	return &userService{
		userRep: userRep,
	}
}

func (u *userService) Create(name string) (model.UserId, error) {
	return u.userRep.Create(name)
}

func (u *userService) Get(id model.UserId) (model.User, error) {
	return u.userRep.Get(id)
}
