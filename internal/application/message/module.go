package message_application

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewDirectMessageUsecase),
	fx.Provide(NewGetMessageHistoryUsecase),
	fx.Provide(NewChatMessageAssembler),
	fx.Provide(NewMessageHistoryAssembler),
	fx.Provide(NewSendMessageUsecase),
	fx.Provide(NewReplyProvider),
)
