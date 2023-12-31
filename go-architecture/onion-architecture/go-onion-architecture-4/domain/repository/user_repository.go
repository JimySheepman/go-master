package repository

import (
	"context"

	"auth/domain/model"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
	GetUser(ctx context.Context, id, email string) (*model.User, error)
}
