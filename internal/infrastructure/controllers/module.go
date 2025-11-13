package controllers

import (
	"go.uber.org/fx"

	auth "main/internal/infrastructure/controllers/auth"
	chat "main/internal/infrastructure/controllers/chat"
	message "main/internal/infrastructure/controllers/message"
	user "main/internal/infrastructure/controllers/user"
	websocket "main/internal/infrastructure/controllers/websocket"
)

var Module = fx.Options(
	auth.Module,
	chat.Module,
	message.Module,
	user.Module,
	websocket.Module,
)
