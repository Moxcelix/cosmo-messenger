package message_application

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDirectMessageUsecase),
	fx.Provide(NewGetMessageHistoryUsecase),
	fx.Provide(NewGetDirectMessageHistoryUsecase),
	fx.Provide(NewSendMessageUsecase),

	fx.Provide(NewChatMessageAssembler),
	fx.Provide(NewMessageHistoryAssembler),
	fx.Provide(NewReplyProvider),
	fx.Provide(NewMessageSender),
)
