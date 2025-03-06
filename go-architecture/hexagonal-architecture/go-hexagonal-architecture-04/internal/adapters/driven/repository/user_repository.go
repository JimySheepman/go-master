package repository

import (
	"context"
	"go-hexagonal/internal/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, sqluser domain.User) error
}
