package userservice_api

import (
	"main/pkg"
)

type UserServiceRoutes struct {
	handler                pkg.RequestHandler
	userRegisterController *UserRegisterController
	userGetInfoController  *UserGetInfoController
}

func NewUserServiceRoutes(
	userRegisterController *UserRegisterController,
	userGetInfoController *UserGetInfoController,
	handler pkg.RequestHandler,
) *UserServiceRoutes {
	return &UserServiceRoutes{
		userGetInfoController:  userGetInfoController,
		userRegisterController: userRegisterController,
		handler:                handler,
	}
}

func (r *UserServiceRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/users")

	group.POST("/register", r.userRegisterController.Register)
	group.GET("/get_info", r.userGetInfoController.GetInfo)
}
