package userservice_application

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRegisterUseCase),
	fx.Provide(NewGetInfoUseCase),
	fx.Provide(NewDeleteUserUsecase),
	fx.Provide(NewGetUsersListUsecase),
)
