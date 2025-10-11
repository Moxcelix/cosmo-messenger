package auth_api

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewLoginController),
	fx.Provide(NewRefreshController),
	fx.Provide(NewValidateController),
	fx.Provide(NewAuthMiddleware),
	fx.Provide(NewAdminAuthMiddleware),
	fx.Provide(NewAuthServiceRoutes),
)
