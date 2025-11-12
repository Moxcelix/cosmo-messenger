package message_application

import (
	"main/internal/application/message/mappers"
	"main/internal/application/message/services"
	"main/internal/application/message/usecases"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(usecases.NewDirectMessageUsecase),
	fx.Provide(usecases.NewGetMessageHistoryUsecase),
	fx.Provide(usecases.NewGetDirectMessageHistoryUsecase),
	fx.Provide(usecases.NewSendMessageUsecase),

	fx.Provide(mappers.NewChatMessageAssembler),
	fx.Provide(mappers.NewMessageHistoryAssembler),
	fx.Provide(mappers.NewReplyProvider),
	fx.Provide(services.NewMessageSender),
)
