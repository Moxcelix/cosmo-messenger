package userservice

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewPasswordHasher),
	fx.Provide(NewDefaultDeleteUserPolicy),
)
