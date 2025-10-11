package auth_application

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewLoginUsecase),
	fx.Provide(NewRefreshUsecase),
	fx.Provide(NewValidateUsecase),
)
