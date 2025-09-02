package bootstrap

import (
	"main/internal/api"
	ping_api "main/internal/api/ping"
	"main/internal/config"
	"main/pkg"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	api.Module,
	config.Module,
	pkg.Module,
	ping_api.Module,
)
