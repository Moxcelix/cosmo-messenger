package message_api

import (
	message_application "main/internal/application/message"
	"main/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendMessageController struct {
	sendMessageUsecase *message_application.SendMessageUsecase
	logger             pkg.Logger
}

func NewSendMessageController(
	sendMessageUsecase *message_application.SendMessageUsecase,
	logger pkg.Logger) *SendMessageController {
	return &SendMessageController{
		sendMessageUsecase: sendMessageUsecase,
		logger:             logger,
	}
}

type sendRequest struct {
	Content string `json:"content" binding:"required"`
}

type sendResponse struct {
	Message string `json:"message"`
}

// SendMessage godoc
// @Summary Message sending
// @Description Send message
// @Tags messages
// @Accept json
// @Produce json
// @Param chat_id path string true "Chat ID"
// @Param input body sendRequest true "Message data"
// @Success 201 {object} sendResponse
// @Failure 400 {object} map[string]string
// @Security     BearerAuth
// @Router /api/v1/messages/chat/{chat_id} [post]
func (c *SendMessageController) SendMessage(ctx *gin.Context) {
	userId := ctx.GetString("UserID")
	chatId := ctx.Param("chat_id")

	var req sendRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.sendMessageUsecase.Execute(userId, chatId, req.Content); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sendResponse{
		Message: "message sent",
	})

}
