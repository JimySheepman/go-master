//go:generate mockgen -package service -source=service.go -destination service_mock.go

// Package service holds service layer methods and repository/database abstractions.
package service

import (
	"context"
	"fmt"
	"time"

	"golang_startup/internal/model"
	"golang_startup/internal/model/request"
)

// Repository abstracts the data access layer.
type Repository interface {
	FindUser(ctx context.Context, id string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (string, error)
	UpdateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, id string) error
	FindAllUsers(ctx context.Context, offset int64, limit int64) ([]model.User, error)
}

// Timer abstraction layer for time operations.
type Timer interface {
	Now() time.Time
}

// Service represents the user domain service layer.
type Service struct {
	repo  Repository
	timer Timer
}

// New instantiates a Service
func New(repo Repository, timer Timer) *Service {
	return &Service{
		repo:  repo,
		timer: timer,
	}
}

// FindUser - Returns a new user for a given ID. Returns the user or an error if anything fails.
func (s *Service) FindUser(ctx context.Context, id string) (model.User, error) {
	user, err := s.repo.FindUser(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("fail to get user: %v", err)
	}

	return user, nil
}

// CreateUser - Creates a new user. Returns the created user or an error if anything fails.
func (s *Service) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	user.CreatedAt = s.timer.Now()
	user.UpdatedAt = s.timer.Now()

	userID, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return model.User{}, fmt.Errorf("fail creating user: %v", err)
	}

	createdUser, err := s.repo.FindUser(ctx, userID)
	if err != nil {
		return model.User{}, fmt.Errorf("fail finding created user: %v", err)
	}

	return createdUser, nil
}

// UpdateUser - Updates an existing user. It returns the updated user or returns an error if anything fails.
func (s *Service) UpdateUser(ctx context.Context, id string, userUpdate request.UserUpdate) (model.User, error) {
	user, err := s.repo.FindUser(ctx, id)
	if err != nil {
		return model.User{}, fmt.Errorf("fail updating user: %v", err)
	}

	user.Name = userUpdate.Name
	user.Email = userUpdate.Email
	user.UpdatedAt = s.timer.Now()

	err = s.repo.UpdateUser(ctx, user)
	if err != nil {
		return model.User{}, fmt.Errorf("fail updating user: %v", err)
	}

	return user, nil
}

// DeleteUser - Removes a user from the database. It returns an error if anything fails.
func (s *Service) DeleteUser(ctx context.Context, id string) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("fail deleting user: %v", err)
	}

	return nil
}

// FindAllUsers - Returns all users, for a given limit and offset
func (s *Service) FindAllUsers(ctx context.Context, offset int64, limit int64) ([]model.User, error) {
	const (
		maxLimit  int64 = 20
		minOffset int64 = 0
	)

	finalLimit := limit
	if finalLimit > maxLimit {
		finalLimit = maxLimit
	}

	finalOffset := offset
	if finalOffset < minOffset {
		finalOffset = minOffset
	}

	users, err := s.repo.FindAllUsers(ctx, finalOffset, finalLimit)
	if err != nil {
		return []model.User{}, fmt.Errorf("fail to get users: %v", err)
	}

	return users, nil
}
