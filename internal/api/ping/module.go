package ping_api

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewPingController),
	fx.Provide(NewPingRoutes),
)
