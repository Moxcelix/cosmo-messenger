package auth_api

import (
	controllers "main/internal/infrastructure/controllers/auth"
	"main/internal/infrastructure/middlewares"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(controllers.NewLoginController),
	fx.Provide(controllers.NewRefreshController),
	fx.Provide(controllers.NewValidateController),

	fx.Provide(middlewares.NewAuthMiddleware),
	fx.Provide(middlewares.NewAdminAuthMiddleware),
	fx.Provide(middlewares.NewQueryAuthMiddleware),

	fx.Provide(NewAuthServiceRoutes),
)
