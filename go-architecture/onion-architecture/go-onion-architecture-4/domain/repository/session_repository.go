package repository

import (
	"context"

	"auth/domain/model"
)

type SessionRepository interface {
	GetSessionByAccessToken(ctx context.Context, accessToken string) (*model.Session, error)
	InsertSession(ctx context.Context, session *model.Session) (*model.Session, error)
	UpdateSession(ctx context.Context, session *model.Session) (*model.Session, error)
	DeleteSession(ctx context.Context, sessionToken string) error
}
