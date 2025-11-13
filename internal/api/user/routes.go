package user_api

import (
	controllers "main/internal/infrastructure/controllers/user"
	"main/internal/infrastructure/middlewares"
	"main/pkg"
)

type UserServiceRoutes struct {
	handler                pkg.RequestHandler
	userRegisterController *controllers.UserRegisterController
	userGetInfoController  *controllers.UserGetInfoController
	userDeleteController   *controllers.UserDeleteController
	getUsersListController *controllers.GetUsersListController
	findUserController     *controllers.FindUserController
	authMiddleware         *middlewares.AuthMiddleware
	adminAuthMiddleware    *middlewares.AdminAuthMiddleware
}

func NewUserServiceRoutes(
	userRegisterController *controllers.UserRegisterController,
	userGetInfoController *controllers.UserGetInfoController,
	userDeleteController *controllers.UserDeleteController,
	getUsersListController *controllers.GetUsersListController,
	findUserController *controllers.FindUserController,
	authMiddleware *middlewares.AuthMiddleware,
	adminAuthMiddleware *middlewares.AdminAuthMiddleware,
	handler pkg.RequestHandler,
) *UserServiceRoutes {
	return &UserServiceRoutes{
		userGetInfoController:  userGetInfoController,
		userRegisterController: userRegisterController,
		userDeleteController:   userDeleteController,
		getUsersListController: getUsersListController,
		findUserController:     findUserController,
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
		authGroup.GET("/find", r.findUserController.FindUser)
	}

	adminGroup := base.Group("")
	adminGroup.Use(r.adminAuthMiddleware.Handler())
	{
		adminGroup.DELETE("/delete/:user_id", r.userDeleteController.DeleteByPath)
	}
}
