package chat_api

import (
	controllers "main/internal/infrastructure/controllers/chat"
	websocket "main/internal/infrastructure/controllers/websocket/chat"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(controllers.NewGetUserChatsController),
	fx.Provide(controllers.NewTypingController),

	fx.Provide(websocket.NewTypingWebSocket),

	fx.Provide(NewChatRoutes),
	fx.Provide(NewChatEvents),
)
