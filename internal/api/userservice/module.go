package userservice_api

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserRegisterController),
	fx.Provide(NewUserServiceRoutes),
)
