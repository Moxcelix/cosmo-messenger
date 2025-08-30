package api

import (
	ping_api "messenger/internal/api/ping"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	pingRoutes ping_api.PingRoutes,
) Routes {
	return Routes{
		pingRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
