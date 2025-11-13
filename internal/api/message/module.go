package message_api

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewMessageRoutes),
	fx.Provide(NewMessageEvents),
)
