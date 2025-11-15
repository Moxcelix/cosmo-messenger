package infrastructure

import (
	"main/internal/infrastructure/controllers"
	"main/internal/infrastructure/middlewares"
	"main/internal/infrastructure/persistence"
	"main/internal/infrastructure/services"
	"main/internal/infrastructure/websocket"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	websocket.Module,
	middlewares.Module,
	persistence.Module,
	services.Module,

	fx.Provide(NewWorkers),
)
