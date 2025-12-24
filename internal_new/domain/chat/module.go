package chat

import (
	"main/internal_new/domain/chat/services"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(services.NewChatHeaderService),
)
