package userservice_api

import (
	"main/internal/api/authservice"
	"main/pkg"
)

type UserServiceRoutes struct {
	handler                pkg.RequestHandler
	userRegisterController *UserRegisterController
	userGetInfoController  *UserGetInfoController
	userDeleteController   *UserDeleteController
	authMiddleware         *authservice_api.AuthMiddleware
}

func NewUserServiceRoutes(
	userRegisterController *UserRegisterController,
	userGetInfoController *UserGetInfoController,
	userDeleteController *UserDeleteController,
	authMiddleware *authservice_api.AuthMiddleware,
	handler pkg.RequestHandler,
) *UserServiceRoutes {
	return &UserServiceRoutes{
		userGetInfoController:  userGetInfoController,
		userRegisterController: userRegisterController,
		userDeleteController:   userDeleteController,
		authMiddleware:         authMiddleware,
		handler:                handler,
	}
}

func (r *UserServiceRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/users")

	group.POST("/register", r.userRegisterController.Register)
	group.GET("/get_info", r.userGetInfoController.GetInfo)

	protected := group.Use(r.authMiddleware.Handler())
	protected.DELETE("/delete", r.userDeleteController.Delete).Use()
}
