package chat_application

import (
	"main/internal/application/chat/services"
	"main/internal/application/chat/usecases"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(usecases.NewGetUserChatsUsecase),
	fx.Provide(usecases.NewTypingUsecase),

	fx.Provide(services.NewChatCollectionAssembler),
	fx.Provide(services.NewChatHeaderProvider),
	fx.Provide(services.NewChatItemAssembler),
	fx.Provide(services.NewLastMessageProvider),
	fx.Provide(services.NewChatNamingService),
	fx.Provide(services.NewChatCreator),
)
