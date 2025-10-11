package bootstrap

import (
	"main/internal/api"
	"main/internal/api/auth"
	ping_api "main/internal/api/ping"
	swagger_api "main/internal/api/swagger"
	user_api "main/internal/api/user"
	"main/internal/application/auth"
	user_application "main/internal/application/user"
	"main/internal/config"
	user "main/internal/domain/user"
	"main/internal/infrastructure/auth"
	user_infrastructure "main/internal/infrastructure/user"
	"main/pkg"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	api.Module,
	config.Module,
	pkg.Module,
	ping_api.Module,
	user_infrastructure.Module,
	user_application.Module,
	swagger_api.Module,
	user_api.Module,
	user.Module,
	auth_infrastructure.Module,
	auth_application.Module,
	auth_api.Module,
)
