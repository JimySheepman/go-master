package drivenport

import (
	"context"
	"go-hexagonal/internal/core/domain"
)

type UserDB interface {
	Create(ctx context.Context, user domain.User) error
}
