package websocket_api

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewWebSocketController),
	fx.Provide(NewWebSocketRoutes),
)
