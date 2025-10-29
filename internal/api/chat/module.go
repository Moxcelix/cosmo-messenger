package chat_api

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewGetUserChatsController),
	fx.Provide(NewTypingWebSocket),
	fx.Provide(NewChatRoutes),
	fx.Provide(NewChatEvents),
)
