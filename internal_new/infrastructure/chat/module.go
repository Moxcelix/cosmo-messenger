package chat

import (
	"main/internal_new/infrastructure/chat/repositories"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(repositories.NewPostgresChatRepository),
	fx.Provide(repositories.NewPostgresMessageRepository),
)
