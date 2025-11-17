package chat_application

import (
	"main/internal/application/chat/mappers"
	"main/internal/application/chat/services"
	"main/internal/application/chat/usecases"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(usecases.NewGetUserChatsUsecase),
	fx.Provide(usecases.NewTypingUsecase),

	fx.Provide(services.NewChatRegistryService),

	fx.Provide(mappers.NewChatCollectionMapper),
	fx.Provide(mappers.NewChatItemMapper),
	fx.Provide(mappers.NewLastMessageMapper),
)
