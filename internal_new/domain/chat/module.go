package chat

import (
	"main/internal_new/domain/chat/factories"
	"main/internal_new/domain/chat/services"
	"main/internal_new/domain/chat/usecases"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(factories.NewChatProjectionFactory),
	fx.Provide(factories.NewMessageProjectionFactory),
	fx.Provide(factories.NewUserProjectionFactory),
	fx.Provide(factories.NewCollectionProjectionFactory),
	fx.Provide(factories.NewChatFactory),

	fx.Provide(services.NewDirectChatService),
	fx.Provide(services.NewCompanionService),
	fx.Provide(services.NewUserChatService),

	fx.Provide(usecases.NewChatCollectionUsecase),
	fx.Provide(usecases.NewChatHistoryUsecase),
	fx.Provide(usecases.NewDirectHistoryUsecase),
	fx.Provide(usecases.NewSendChatMessageUsecase),
	fx.Provide(usecases.NewSendDirectMessageUsecase),
)
