package repository

import (
	"context"

	"bastet/internal/repository/sqlc"

	"github.com/jackc/pgx/v4/pgxpool"
)

// repository hides the core.Repository implementation of the database operations
// Needs a sqlc.Querier, which will allow for the communication between the SQLC
// layer (automatically created) and the DB interface we expose to the service
type repository struct {
	querier sqlc.Querier
}

// New creates a new repository given a database URL.
func New(URL string) (*repository, error) {
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, URL)
	if err != nil {
		return nil, err
	}

	querier := sqlc.New(pool)

	repo := &repository{
		querier: querier,
	}
	return repo, nil
}
