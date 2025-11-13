package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAdminAuthMiddleware),
	fx.Provide(NewAuthMiddleware),
	fx.Provide(NewQueryAuthMiddleware),
)
