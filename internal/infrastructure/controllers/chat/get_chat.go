package controllers

import (
	chat_application "main/internal/application/chat/usecases"
	chat_domain "main/internal/domain/chat"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetChatController struct {
	chatUsecase *chat_application.GetChatUsecase
}

func NewGetChatController(
	chatUsecase *chat_application.GetChatUsecase,
) *GetChatController {
	return &GetChatController{
		chatUsecase: chatUsecase,
	}
}

// GetChat godoc
// @Summary Get specific chat
// @Description Get detailed information about a specific chat
// @Tags chats
// @Accept json
// @Produce json
// @Param chat_id path string true "Chat ID"
// @Security BearerAuth
// @Success 200 {object} dto.ChatItem
// @Failure 400 {object} map[string]string "Invalid chat ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Access denied"
// @Failure 404 {object} map[string]string "Chat not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/chats/{chat_id} [get]
func (c *GetChatController) GetChat(ctx *gin.Context) {
	userID := ctx.GetString("UserID")
	chatID := ctx.Param("chat_id")

	if chatID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "chat ID is required"})
		return
	}

	chat, err := c.chatUsecase.Execute(userID, chatID)
	if err != nil {
		switch err {
		case chat_domain.ErrChatNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": "chat not found"})
		case chat_domain.ErrChatAccessDenied:
			ctx.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, chat)
}
