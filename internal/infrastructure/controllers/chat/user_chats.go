package controllers

import (
	chat_application "main/internal/application/chat/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUserChatsController struct {
	userChatsUsecase *chat_application.GetUserChatsUsecase
}

func NewGetUserChatsController(
	userChatsUsecase *chat_application.GetUserChatsUsecase,
) *GetUserChatsController {
	return &GetUserChatsController{
		userChatsUsecase: userChatsUsecase,
	}
}

// GetUserChats godoc
// @Summary Get user chats
// @Description Get paginated list of user's chats with cursor-based pagination
// @Tags chats
// @Accept json
// @Produce json
// @Param cursor query string false "Cursor for pagination (chat ID)"
// @Param count query int false "Number of chats per page" default(10)
// @Param direction query string false "Pagination direction: older or newer" Enums(older, newer) default(older)
// @Security BearerAuth
// @Success 200 {object} dto.ChatCollection
// @Failure 400 {object} map[string]string "Invalid parameters"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/chats [get]
func (c *GetUserChatsController) GetUserChats(ctx *gin.Context) {
	userID := ctx.GetString("UserID")

	cursorChatID := ctx.Query("cursor")

	count, err := strconv.Atoi(ctx.DefaultQuery("count", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid count parameter"})
		return
	}

	direction := ctx.DefaultQuery("direction", "older")
	if direction != "older" && direction != "newer" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "direction must be 'older' or 'newer'"})
		return
	}

	list, err := c.userChatsUsecase.Execute(userID, cursorChatID, count, direction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, list)
}
