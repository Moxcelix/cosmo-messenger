package chat_infrastructure

import (
	chat_domain "main/internal/domain/chat"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewChatPolicyConfig),
	fx.Provide(NewChatRepository),
	fx.Provide(NewWebsocketTypingBroadcaster),
	fx.Provide(NewInMemoryTypingService),
	fx.Provide(
		fx.Annotate(
			func(service *InMemoryTypingService) chat_domain.TypingService {
				return service
			},
		),
	),
)
