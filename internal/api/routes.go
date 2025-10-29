package api

import (
	auth_api "main/internal/api/auth"
	chat_api "main/internal/api/chat"
	message_api "main/internal/api/message"
	ping_api "main/internal/api/ping"
	swagger_api "main/internal/api/swagger"
	user_api "main/internal/api/user"
	websocket_api "main/internal/api/websocket"
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	swaggerRoutes *swagger_api.SwaggerRoutes,
	pingRoutes *ping_api.PingRoutes,
	userRoutes *user_api.UserServiceRoutes,
	authRoutes *auth_api.AuthServiceRoutes,
	msgRoutes *message_api.MessageRoutes,
	chatRoutes *chat_api.ChatRoutes,
	websocketRoutes *websocket_api.WebSocketRoutes,
) Routes {
	return Routes{
		swaggerRoutes,
		pingRoutes,
		userRoutes,
		authRoutes,
		msgRoutes,
		chatRoutes,
		websocketRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
