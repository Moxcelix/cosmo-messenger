package auth_application

import (
	"main/internal/application/auth/usecases"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(usecases.NewLoginUsecase),
	fx.Provide(usecases.NewRefreshUsecase),
	fx.Provide(usecases.NewValidateUsecase),
)
