package message_api

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDirectMessageController),
	fx.Provide(NewGetChatMessagesController),
	fx.Provide(NewMessageRoutes),
)
