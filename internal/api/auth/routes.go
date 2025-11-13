package auth_api

import (
	controllers "main/internal/infrastructure/controllers/auth"
	"main/pkg"
)

type AuthServiceRoutes struct {
	handler            pkg.RequestHandler
	loginController    *controllers.LoginController
	refreshController  *controllers.RefreshController
	validateController *controllers.ValidateController
}

func NewAuthServiceRoutes(
	handler pkg.RequestHandler,
	loginController *controllers.LoginController,
	refreshController *controllers.RefreshController,
	validateController *controllers.ValidateController) *AuthServiceRoutes {
	return &AuthServiceRoutes{
		handler:            handler,
		validateController: validateController,
		loginController:    loginController,
		refreshController:  refreshController,
	}
}

func (r *AuthServiceRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/auth")

	group.POST("/login", r.loginController.Login)
	group.POST("/refresh", r.refreshController.Refresh)
	group.POST("/validate", r.validateController.Validate)
}
