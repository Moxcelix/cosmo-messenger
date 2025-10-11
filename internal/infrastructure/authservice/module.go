package authservice_infrastructure

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewInternalAuthService),
)
