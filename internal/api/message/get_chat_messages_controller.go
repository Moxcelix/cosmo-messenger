package message_api

import (
	message_application "main/internal/application/message"
	message_domain "main/internal/domain/message"
	"main/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetChatMessagesController struct {
	getChatMessagesUsecase *message_application.GetChatMessagesUsecase
	logger                 pkg.Logger
}

func NewGetChatMessagesController(
	getChatMessagesUsecase *message_application.GetChatMessagesUsecase,
	logger pkg.Logger,
) *GetChatMessagesController {
	return &GetChatMessagesController{
		getChatMessagesUsecase: getChatMessagesUsecase,
		logger:                 logger,
	}
}

type chatMessagesResponse struct {
	Data []messageData  `json:"data"`
	Meta paginationMeta `json:"meta"`
}

type messageData struct {
	message_domain.Message
}

type paginationMeta struct {
	HasPrev    bool `json:"has_prev"`
	HasNext    bool `json:"has_next"`
	Page       int  `json:"page"`
	TotalPages int  `json:"total_pages"`
	Total      int  `json:"total"`
	Count      int  `json:"count"`
}

// GetChatMessages godoc
// @Summary Get chat messages
// @Description Get paginated messages from a chat
// @Tags messages
// @Accept json
// @Produce json
// @Param chat_id path string true "Chat ID"
// @Param page query int false "Page number" default(1)
// @Param count query int false "Number of messages per page" default(10)
// @Security BearerAuth
// @Success 200 {object} chatMessagesResponse
// @Failure 400 {object} map[string]string "Invalid parameters"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Access denied to chat"
// @Failure 404 {object} map[string]string "Chat not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/messages/chat/{chat_id} [get]
func (c *GetChatMessagesController) GetChatMessages(ctx *gin.Context) {
	userId := ctx.GetString("UserID")
	chatId := ctx.Param("chat_id")

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid page parameter"})
		return
	}

	count, err := strconv.Atoi(ctx.DefaultQuery("count", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid count parameter"})
		return
	}

	messages, err := c.getChatMessagesUsecase.Execute(userId, chatId, page, count)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	messageDatas := make([]messageData, len(messages.Messages))
	for i, msg := range messages.Messages {
		messageDatas[i] = messageData{*msg}
	}

	totalPages := (messages.Total + count - 1) / count
	if totalPages == 0 {
		totalPages = 1
	}

	response := chatMessagesResponse{
		Data: messageDatas,
		Meta: paginationMeta{
			HasPrev:    page > 1,
			HasNext:    page < totalPages,
			Page:       page,
			TotalPages: totalPages,
			Total:      messages.Total,
			Count:      len(messageDatas),
		},
	}

	ctx.JSON(http.StatusOK, response)
}
