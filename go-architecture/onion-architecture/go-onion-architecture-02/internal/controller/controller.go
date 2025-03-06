//go:generate mockgen -package controller -source=controller.go -destination controller_mock.go

// Package controller holds application controllers.
package controller

import (
	"context"

	"golang_startup/internal/model"
	"golang_startup/internal/model/request"

	"github.com/gorilla/mux"
)

// Service abstracts the service layer.
type Service interface {
	FindUser(ctx context.Context, id string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUser(ctx context.Context, id string, userUpdate request.UserUpdate) (model.User, error)
	DeleteUser(ctx context.Context, id string) error
	FindAllUsers(ctx context.Context, offset int64, limit int64) ([]model.User, error)
}

// Controller - controller related struct
type Controller struct {
	mux     *mux.Router
	service Service
}

// New - Instantiates a new Controller. Requires a Service.
func New(service Service) *Controller {
	c := &Controller{
		mux:     mux.NewRouter(),
		service: service,
	}

	c.routing()

	return c
}

// Mux - Initializes routing and returns a Router.
func (c *Controller) Mux() *mux.Router {
	c.routing()

	return c.mux
}
