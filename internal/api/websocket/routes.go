package websocket_api

import (
	controllers "main/internal/infrastructure/controllers/websocket"
	"main/internal/infrastructure/middlewares"
	"main/pkg"
)

type WebSocketRoutes struct {
	handler             pkg.RequestHandler
	websocketController *controllers.WebSocketController
	authMiddleware      *middlewares.QueryAuthMiddleware
}

func NewWebSocketRoutes(
	websocketController *controllers.WebSocketController,
	authMiddleware *middlewares.QueryAuthMiddleware,
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
