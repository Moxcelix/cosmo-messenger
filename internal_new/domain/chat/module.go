package chat

import (
	"main/internal_new/domain/chat/services"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(services.NewChatHeaderService),
	fx.Provide(services.NewMessageDemoService),
	fx.Provide(services.NewMessageDefaultService),
	fx.Provide(services.NewMessageReplyService),
	fx.Provide(services.NewSenderService),
)
