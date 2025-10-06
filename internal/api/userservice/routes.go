package userservice_api

import (
	"main/pkg"
)

type UserServiceRoutes struct {
	handler                pkg.RequestHandler
	userRegisterController *UserRegisterController
	userGetInfoController  *UserGetInfoController
	userDeleteController   *UserDeleteController
}

func NewUserServiceRoutes(
	userRegisterController *UserRegisterController,
	userGetInfoController *UserGetInfoController,
	userDeleteController *UserDeleteController,
	handler pkg.RequestHandler,
) *UserServiceRoutes {
	return &UserServiceRoutes{
		userGetInfoController:  userGetInfoController,
		userRegisterController: userRegisterController,
		userDeleteController:   userDeleteController,
		handler:                handler,
	}
}

func (r *UserServiceRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/users")

	group.POST("/register", r.userRegisterController.Register)
	group.GET("/get_info", r.userGetInfoController.GetInfo)
	group.DELETE("/delete", r.userDeleteController.Delete)
}
