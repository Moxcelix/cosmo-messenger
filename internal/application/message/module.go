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

	fx.Provide(mappers.NewAttachmentMapper),
	fx.Provide(mappers.NewChatHeaderMapper),
	fx.Provide(mappers.NewReplyMapper),
	fx.Provide(mappers.NewMessageMapper),
	fx.Provide(mappers.NewMessageHistoryMapper),

	fx.Provide(services.NewMessageDispatcher),
	fx.Provide(services.NewMessageHistoryService),
)
