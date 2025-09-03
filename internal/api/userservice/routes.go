package userservice_api

import (
	"main/pkg"
)

type UserServiceRoutes struct {
	handler                pkg.RequestHandler
	userRegisterController *UserRegisterController
}

func NewUserServiceRoutes(
	userRegisterRegister *UserRegisterController,
	handler pkg.RequestHandler,
) *UserServiceRoutes {
	return &UserServiceRoutes{
		userRegisterController: userRegisterRegister,
		handler:                handler,
	}
}

func (r *UserServiceRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/users")

	group.POST("/register", r.userRegisterController.Register)
}
