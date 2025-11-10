package message_api

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDirectMessageController),
	fx.Provide(NewGetChatMessagesController),
	fx.Provide(NewGetDirectMessagesController),
	fx.Provide(NewSendMessageController),
	fx.Provide(NewSendMessageWebSocket),
	fx.Provide(NewMessageRoutes),
	fx.Provide(NewMessageEvents),
)
