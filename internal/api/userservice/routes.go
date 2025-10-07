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
	adminAuthMiddleware    *authservice_api.AdminAuthMiddleware
}

func NewUserServiceRoutes(
	userRegisterController *UserRegisterController,
	userGetInfoController *UserGetInfoController,
	userDeleteController *UserDeleteController,
	authMiddleware *authservice_api.AuthMiddleware,
	adminAuthMiddleware *authservice_api.AdminAuthMiddleware,
	handler pkg.RequestHandler,
) *UserServiceRoutes {
	return &UserServiceRoutes{
		userGetInfoController:  userGetInfoController,
		userRegisterController: userRegisterController,
		userDeleteController:   userDeleteController,
		authMiddleware:         authMiddleware,
		adminAuthMiddleware:    adminAuthMiddleware,
		handler:                handler,
	}
}

func (r *UserServiceRoutes) Setup() {
	base := r.handler.Gin.Group("/api/v1/users")

	baseGroup := base.Group("")
	{
		baseGroup.POST("/register", r.userRegisterController.Register)
		baseGroup.GET("/get_info", r.userGetInfoController.GetInfo)
	}

	authGroup := base.Group("")
	authGroup.Use(r.authMiddleware.Handler())
	{
		authGroup.DELETE("/delete", r.userDeleteController.Delete)
	}

	adminGroup := base.Group("/admin")
	adminGroup.Use(r.adminAuthMiddleware.Handler())
	{
		adminGroup.DELETE("/delete/:username", r.userDeleteController.Delete)
	}
}
