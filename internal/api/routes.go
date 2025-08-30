package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RegisterRoutes(lc fx.Lifecycle, ping *PingController) {
	r := gin.Default()

	r.GET("/ping", ping.Ping)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go r.Run(":5005")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}

var Module = fx.Options(
	fx.Provide(NewPingController),
	fx.Invoke(RegisterRoutes),
)
