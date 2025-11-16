package persistence

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewChatPolicyConfig),
	fx.Provide(NewChatRepository),
	fx.Provide(NewChatListQuery),
	fx.Provide(NewChatQuery),
)
