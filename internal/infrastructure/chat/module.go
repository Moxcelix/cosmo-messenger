package chat_infrastructure

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewChatPolicyConfig),
	fx.Provide(NewChatRepository),
	fx.Provide(NewWebsocketTypingBroadcaster),
	fx.Provide(NewWebsocketChatBroadcaster),
)
