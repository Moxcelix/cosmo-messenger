package bootstrap

import (
	"go.uber.org/fx"
	"messenger/internal/api"
)

var CommonModules = fx.Options(
	api.Module,
)
