package controllers

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewLoginController),
	fx.Provide(NewRefreshController),
	fx.Provide(NewValidateController),
)
