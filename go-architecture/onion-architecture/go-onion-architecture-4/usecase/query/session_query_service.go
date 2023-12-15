package query

import (
	"context"

	"auth/domain/model"
)

type SessionQueryService interface {
	GetSessionBySessionToken(ctx context.Context, sessionToken string) (*model.Session, error)
}
