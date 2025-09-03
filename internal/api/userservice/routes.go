package userservice_api

import (
	"main/pkg"
)

type UserServiceRoutes struct {
	handler                pkg.RequestHandler
	userRegisterController *UserRegisterController
}

func NewUserServiceRoutes(
	userRegisterController *UserRegisterController,
	handler pkg.RequestHandler,
) *UserServiceRoutes {
	return &UserServiceRoutes{
		userRegisterController: userRegisterController,
		handler:                handler,
	}
}

func (r *UserServiceRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/users")

	group.POST("/register", r.userRegisterController.Register)
}
