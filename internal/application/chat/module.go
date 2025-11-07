package chat_application

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewGetUserChatsUsecase),
	fx.Provide(NewTypingUsecase),

	fx.Provide(NewChatCollectionAssembler),
	fx.Provide(NewChatHeaderProvider),
	fx.Provide(NewChatItemAssembler),
	fx.Provide(NewLastMessageProvider),
	fx.Provide(NewChatNamingService),
)
