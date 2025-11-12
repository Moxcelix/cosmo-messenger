package websocket_api

import (
	message_application "main/internal/application/message/usecases"
	"main/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebSocketController struct {
	sendMessageUsecase *message_application.SendMessageUsecase
	logger             pkg.Logger
	wsHub              *pkg.WebSocketHub
}

func NewWebSocketController(
	sendMessageUsecase *message_application.SendMessageUsecase,
	logger pkg.Logger,
	wsHub *pkg.WebSocketHub,
) *WebSocketController {
	return &WebSocketController{
		sendMessageUsecase: sendMessageUsecase,
		logger:             logger,
		wsHub:              wsHub,
	}
}

// HandleWebSocket godoc
// @Summary WebSocket connection for real-time messages
// @Description WebSocket endpoint for real-time messaging
// @Tags websocket
// @Router /ws [get]
func (c *WebSocketController) HandleWebSocket(ctx *gin.Context) {
	userID := ctx.GetString("UserID")
	if userID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user_id required"})
		return
	}

	err := c.wsHub.HandleConnection(ctx.Writer, ctx.Request, userID)
	if err != nil {
		c.logger.Error("WebSocket connection failed:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket connection failed"})
		return
	}

	c.logger.Info("WebSocket client connected via hub:", userID)
}
