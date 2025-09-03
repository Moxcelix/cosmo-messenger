package bootstrap

import (
	"main/internal/api"
	ping_api "main/internal/api/ping"
	userservice_application "main/internal/application/userservice"
	"main/internal/config"
	userservice_infrastructure "main/internal/infrastructure/userservice"
	"main/pkg"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	api.Module,
	config.Module,
	pkg.Module,
	ping_api.Module,
	userservice_infrastructure.Module,
	userservice_application.Module,
)
