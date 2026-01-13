package auth

import (
	"main/internal_new/infrastructure/auth/repositories"
	"main/internal_new/infrastructure/auth/services"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(services.NewInternalAuthService),
	fx.Provide(services.NewPasswordHasher),

	fx.Provide(repositories.NewPostgresUserRepository),
)
