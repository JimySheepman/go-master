package repository

import (
	"context"

	"auth/domain/model"
)

type AccountRepository interface {
	InsertAccount(ctx context.Context, account *model.Account) (*model.Account, error)
}
