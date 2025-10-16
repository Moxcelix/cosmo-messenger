package chat_application

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewGetUserChatsUsecase),
)
