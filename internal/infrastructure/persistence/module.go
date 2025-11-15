package persistence

import (
	chat "main/internal/infrastructure/persistence/chat"
	message "main/internal/infrastructure/persistence/message"
	user "main/internal/infrastructure/persistence/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	chat.Module,
	message.Module,
	user.Module,
)
