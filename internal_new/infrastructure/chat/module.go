package chat

import (
	"main/internal_new/infrastructure/chat/repositories"
	"main/internal_new/infrastructure/chat/services"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(repositories.NewPostgresChatRepository),
	fx.Provide(repositories.NewPostgresMessageRepository),

	fx.Provide(services.NewMessageWebsocketPublisher),
)
