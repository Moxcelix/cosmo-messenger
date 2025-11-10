package websocket_api

import (
	auth_api "main/internal/api/auth"
	"main/pkg"
)

type WebSocketRoutes struct {
	handler             pkg.RequestHandler
	websocketController *WebSocketController
	authMiddleware      *auth_api.QueryAuthMiddleware
}

func NewWebSocketRoutes(
	websocketController *WebSocketController,
	authMiddleware *auth_api.QueryAuthMiddleware,
	handler pkg.RequestHandler,
) *WebSocketRoutes {
	return &WebSocketRoutes{
		websocketController: websocketController,
		authMiddleware:      authMiddleware,
		handler:             handler,
	}
}

func (r *WebSocketRoutes) Setup() {
	base := r.handler.Gin.Group("/ws").Use(r.authMiddleware.Handler())

	base.GET("/", r.websocketController.HandleWebSocket)
}
