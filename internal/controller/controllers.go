package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/matisiekpl/gobackend/internal/service"
)

type Controllers interface {
	User() UserController

	Route(e *echo.Echo)
}

type controllers struct {
	userController UserController
}

func NewControllers(services service.Services) Controllers {
	userController := newUserController(services.User())
	return &controllers{
		userController: userController,
	}

}

func (c controllers) User() UserController {
	return c.userController
}

func (c controllers) Route(e *echo.Echo) {
	e.POST("/api/auth/login", c.userController.Login)
	e.POST("/api/auth/register", c.userController.Register)
}
