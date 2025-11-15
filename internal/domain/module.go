package domain

import (
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	user_domain "main/internal/domain/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	user_domain.Module,
	message_domain.Module,
	chat_domain.Module,
)
