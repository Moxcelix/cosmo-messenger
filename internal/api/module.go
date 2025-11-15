package api

import (
	auth_api "main/internal/api/auth"
	chat_api "main/internal/api/chat"
	message_api "main/internal/api/message"
	ping_api "main/internal/api/ping"
	swagger_api "main/internal/api/swagger"
	user_api "main/internal/api/user"
	websocket_api "main/internal/api/websocket"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewEvents),

	auth_api.Module,
	chat_api.Module,
	message_api.Module,
	ping_api.Module,
	swagger_api.Module,
	user_api.Module,
	websocket_api.Module,
)
