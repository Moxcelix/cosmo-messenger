package persistence

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewMessagePolicyConfig),
	fx.Provide(NewMessageRepository),
)
