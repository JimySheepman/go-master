package handler

import (
	"go-onion-architecture-sample/common"
	"go-onion-architecture-sample/registory"
	"go-onion-architecture-sample/userinterface/request"
	"go-onion-architecture-sample/userinterface/response"

	"github.com/labstack/echo/v4"
)

func UserHandler(client *echo.Echo, registory registory.Registory) {
	CreateUserHandler(client, registory)
	GetUserHandler(client, registory)
}

func CreateUserHandler(client *echo.Echo, registory registory.Registory) {
	client.POST("/users", func(c echo.Context) error {
		var r request.CreateUserRequest
		if err := c.Bind(&r); err != nil {
			return err
		}

		id, err := registory.UserService().Create(r.Name)
		if err != nil {
			return err
		}

		return c.JSON(200, response.CreateUserResponse{Id: id})
	})
}

func GetUserHandler(client *echo.Echo, registory registory.Registory) {
	client.GET("/user/:id", func(c echo.Context) error {
		id, err := common.GetUserId(c.Param("id"))
		if err != nil {
			return err
		}

		user, err := registory.UserService().Get(id)
		if err != nil {
			return err
		}

		return c.JSON(200, response.GetUserResponse{User: user})
	})
}
