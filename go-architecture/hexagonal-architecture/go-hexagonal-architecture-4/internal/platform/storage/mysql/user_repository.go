package mysqldb

import (
	"context"
	"database/sql"
	"fmt"
	"go-hexagonal/internal/core/domain"
	"time"

	"github.com/huandu/go-sqlbuilder"
)

type UserRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewUserRepository initializes a MySQL-based implementation of UserRepository.
func NewUserRepository(db *sql.DB, dbTimeout time.Duration) *UserRepository {
	return &UserRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the adapter userRepository interface.
func (r *UserRepository) Save(ctx context.Context, user domain.User) error {
	userSQLStruct := sqlbuilder.NewStruct(new(SqlUser))
	query, args := userSQLStruct.InsertInto(sqlUserTable, SqlUser{
		ID:        user.ID.String(),
		Name:      user.Name,
		Lastname:  user.Lastname,
		Email:     user.Email.String(),
		Password:  user.Password.String(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Error trying to persist course on database: %v", err)
	}

	return nil
}
