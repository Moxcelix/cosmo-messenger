package user_application

import (
	"main/internal/application/user/services"
	"main/internal/application/user/usecases"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(usecases.NewRegisterUseCase),
	fx.Provide(usecases.NewGetInfoUseCase),
	fx.Provide(usecases.NewDeleteUserUsecase),
	fx.Provide(usecases.NewGetUsersListUsecase),
	fx.Provide(usecases.NewFindUserUsecase),

	fx.Provide(services.NewSenderProvider),
)
