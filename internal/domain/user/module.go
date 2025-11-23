package user_domain

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewPasswordHasher),
)
