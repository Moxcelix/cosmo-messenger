package bootstrap

import (
	"main/internal/api"
	"main/internal/application"
	"main/internal/config"
	"main/internal/domain"
	"main/internal/infrastructure"
	"main/pkg"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	config.Module,
	pkg.Module,
	domain.Module,
	application.Module,
	infrastructure.Module,
	api.Module,
)
