package user_api

import (
	controllers "main/internal/infrastructure/controllers/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(controllers.NewUserRegisterController),
	fx.Provide(controllers.NewUserGetInfoController),
	fx.Provide(controllers.NewUserDeleteController),
	fx.Provide(controllers.NewGetUsersListController),
	fx.Provide(controllers.NewFindUserController),

	fx.Provide(NewUserServiceRoutes),
)
