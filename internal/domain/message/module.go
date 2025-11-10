package message_domain

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewMessagePolicy),
)
