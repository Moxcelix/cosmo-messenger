package controllers

import (
	message_application "main/internal/application/message/usecases"
	"main/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetHistoryController struct {
	getChatMessagesUsecase *message_application.GetMessageHistoryUsecase
	logger                 pkg.Logger
}

func NewGetHistoryController(
	getChatMessagesUsecase *message_application.GetMessageHistoryUsecase,
	logger pkg.Logger,
) *GetHistoryController {
	return &GetHistoryController{
		getChatMessagesUsecase: getChatMessagesUsecase,
		logger:                 logger,
	}
}

// GetHistory godoc
// @Summary Get chat messages
// @Description Get paginated messages from a chat
// @Tags messages
// @Accept json
// @Produce json
// @Param chat_id path string true "Chat ID"
// @Param cursor query string false "Cursor Message ID"
// @Param dir query string false "Scrolling direction" default("older")
// @Param count query int false "Number of messages per page" default(10)
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string "Invalid parameters"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Access denied to chat"
// @Failure 404 {object} map[string]string "Chat not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/messages/chat/{chat_id} [get]
func (c *GetHistoryController) GetHistory(ctx *gin.Context) {
	userId := ctx.GetString("UserID")
	chatId := ctx.Param("chat_id")

	cursorMessageId := ctx.Query("cursor")
	direction := ctx.Query("dir")

	count, err := strconv.Atoi(ctx.DefaultQuery("count", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid count parameter"})
		return
	}

	messages, err := c.getChatMessagesUsecase.Execute(
		userId, chatId, cursorMessageId, count, direction)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	ctx.JSON(http.StatusOK, messages)
}
