package message_application

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDirectMessageUsecase),
	fx.Provide(NewGetChatMessagesUsecase),
)
