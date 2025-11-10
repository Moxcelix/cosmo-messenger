package chat_api

import (
	chat_application "main/internal/application/chat"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUserChatsController struct {
	userChatsUsecase *chat_application.GetUserChatsUsecase
}

func NewGetUserChatsController(
	userChatsUsecase *chat_application.GetUserChatsUsecase) *GetUserChatsController {
	return &GetUserChatsController{
		userChatsUsecase: userChatsUsecase,
	}
}

// GetUserChats godoc
// @Summary Get user chats
// @Description Get paginated list of user's chats
// @Tags chats
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param count query int false "Number of chats per page" default(10)
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string "Invalid parameters"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/chats [get]
func (c *GetUserChatsController) GetUserChats(ctx *gin.Context) {
	userId := ctx.GetString("UserID")
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

	list, err := c.userChatsUsecase.Execute(userId, page, count)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, list)
}
