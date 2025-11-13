package websocket

import (
	chat "main/internal/infrastructure/websocket/chat"
	message "main/internal/infrastructure/websocket/message"

	"go.uber.org/fx"
)

var Module = fx.Options(
	chat.Module,
	message.Module,
)
