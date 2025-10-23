package chat_application

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewGetUserChatsUsecase),
	fx.Provide(NewChatCollectionAssembler),
	fx.Provide(NewLastMessageProvider),
	fx.Provide(NewChatPolicy),
)
