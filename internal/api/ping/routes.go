package ping_api

import (
	"main/pkg"
)

type PingRoutes struct {
	handler        pkg.RequestHandler
	pingController PingController
}

func NewPingRoutes(
	pingController PingController,
	handler pkg.RequestHandler,
) *PingRoutes {
	return &PingRoutes{
		pingController: pingController,
		handler:        handler,
	}
}

func (r *PingRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1")

	group.GET("/ping", r.pingController.Ping)
}
