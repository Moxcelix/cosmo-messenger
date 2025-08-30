package bootstrap

import (
	"messenger/internal/api"
	ping_api "messenger/internal/api/ping"
	"messenger/internal/config"
	"messenger/pkg"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	api.Module,
	config.Module,
	pkg.Module,
	ping_api.Module,
)
