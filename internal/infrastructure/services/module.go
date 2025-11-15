package services

import (
	auth "main/internal/infrastructure/services/auth"
	chat "main/internal/infrastructure/services/chat"
	message "main/internal/infrastructure/services/message"

	"go.uber.org/fx"
)

var Module = fx.Options(
	auth.Module,
	message.Module,
	chat.Module,
)
