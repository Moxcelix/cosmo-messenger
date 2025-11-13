package infrastructure

import (
	"main/internal/infrastructure/controllers"
	"main/internal/infrastructure/middlewares"
	"main/internal/infrastructure/websocket"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	websocket.Module,
	middlewares.Module,

	fx.Provide(NewWorkers),
)
