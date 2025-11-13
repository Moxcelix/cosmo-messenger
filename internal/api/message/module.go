package message_api

import (
	controllers "main/internal/infrastructure/controllers/message"
	websocket "main/internal/infrastructure/controllers/websocket/message"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(controllers.NewDirectMessageController),
	fx.Provide(controllers.NewGetChatMessagesController),
	fx.Provide(controllers.NewGetDirectMessagesController),
	fx.Provide(controllers.NewSendMessageController),

	fx.Provide(websocket.NewSendMessageWebSocket),

	fx.Provide(NewMessageRoutes),
	fx.Provide(NewMessageEvents),
)
