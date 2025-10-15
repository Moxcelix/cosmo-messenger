package user_api

import (
	auth_api "main/internal/api/auth"
	"main/pkg"
)

type UserServiceRoutes struct {
	handler                pkg.RequestHandler
	userRegisterController *UserRegisterController
	userGetInfoController  *UserGetInfoController
	userDeleteController   *UserDeleteController
	getUsersListController *GetUsersListController
	authMiddleware         *auth_api.AuthMiddleware
	adminAuthMiddleware    *auth_api.AdminAuthMiddleware
}

func NewUserServiceRoutes(
	userRegisterController *UserRegisterController,
	userGetInfoController *UserGetInfoController,
	userDeleteController *UserDeleteController,
	getUsersListController *GetUsersListController,
	authMiddleware *auth_api.AuthMiddleware,
	adminAuthMiddleware *auth_api.AdminAuthMiddleware,
	handler pkg.RequestHandler,
) *UserServiceRoutes {
	return &UserServiceRoutes{
		userGetInfoController:  userGetInfoController,
		userRegisterController: userRegisterController,
		userDeleteController:   userDeleteController,
		getUsersListController: getUsersListController,
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
		baseGroup.GET("/get_usernames_list", r.getUsersListController.GetUsernameList)
	}

	authGroup := base.Group("")
	authGroup.Use(r.authMiddleware.Handler())
	{
		authGroup.DELETE("/delete", r.userDeleteController.DeleteByContext)
	}

	adminGroup := base.Group("")
	adminGroup.Use(r.adminAuthMiddleware.Handler())
	{
		adminGroup.DELETE("/delete/:username", r.userDeleteController.DeleteByPath)
	}
}
