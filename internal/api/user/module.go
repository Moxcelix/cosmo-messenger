package user_api

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserRegisterController),
	fx.Provide(NewUserGetInfoController),
	fx.Provide(NewUserDeleteController),
	fx.Provide(NewGetUsersListController),
	fx.Provide(NewFindUserController),
	fx.Provide(NewUserServiceRoutes),
)
