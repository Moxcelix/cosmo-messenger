package cmd

import (
	"main/bootstrap"

	"context"

	"go.uber.org/fx"

	"main/pkg"

	"main/internal/api"
	"main/internal/config"
	"main/internal/infrastructure"
)

func Run() any {
	return func(
		env config.Env,
		logger pkg.Logger,
		handler pkg.RequestHandler,
		routes api.Routes,
		events api.Events,
		workers infrastructure.Workers,
	) {
		routes.Setup()
		events.Setup()
		workers.Run()
		err := handler.Gin.Run(":" + env.Port)

		if err != nil {
			logger.Error(err)
			return
		}
	}
}

func StartApp() error {
	opts := fx.Options(
		fx.Invoke(Run()),
	)

	app := fx.New(
		bootstrap.CommonModules,
		opts,
	)
	ctx := context.Background()
	err := app.Start(ctx)
	return err
}
