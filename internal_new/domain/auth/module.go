package auth

import (
	"main/internal_new/domain/auth/usecases"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(usecases.NewLoginUsecase),
	fx.Provide(usecases.NewRefreshUsecase),
	fx.Provide(usecases.NewValidateUsecase),
	fx.Provide(usecases.NewRegisterUseCase),
)
