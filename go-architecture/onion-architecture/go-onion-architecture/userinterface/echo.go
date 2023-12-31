package ui

import (
	"go-onion-architecture-sample/registory"
	"go-onion-architecture-sample/userinterface/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ApiClient struct {
	client    *echo.Echo
	registory registory.Registory
}

func NewApiClient(registory registory.Registory) *ApiClient {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &ApiClient{
		client:    e,
		registory: registory,
	}
}

func (a *ApiClient) RegisterRoute() {
	// user
	handler.UserHandler(a.client, a.registory)

	// todo
	handler.TodoHandler(a.client, a.registory)
}

func (a *ApiClient) Start() {
	a.client.Logger.Fatal(a.client.Start(":8080"))
}
