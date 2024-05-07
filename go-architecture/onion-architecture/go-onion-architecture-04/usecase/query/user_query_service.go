package query

import (
	"context"

	"auth/domain/model"
)

type UserQueryService interface {
	FetchUserByProviderAccountID(ctx context.Context, providerID, providerAccountID string) (*model.User, error)
}
