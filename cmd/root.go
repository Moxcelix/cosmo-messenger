package cmd

import (
	"messenger/bootstrap"

	"context"

	"go.uber.org/fx"

	"messenger/pkg"

	route "messenger/internal/api"
	"messenger/internal/config"
)

func Run() any {
	return func(
		env config.Env,
		logger pkg.Logger,
		handler pkg.RequestHandler,
		route route.Routes,
	) {
		route.Setup()
		err := handler.Gin.Run(env.ApiURL + ":" + env.Port)

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
