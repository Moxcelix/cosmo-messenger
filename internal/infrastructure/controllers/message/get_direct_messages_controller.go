package controllers

import (
	"errors"
	message_application "main/internal/application/message/usecases"
	user_domain "main/internal/domain/user"
	"main/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetDirectMessagesController struct {
	usecase *message_application.GetDirectMessageHistoryUsecase
	logger  pkg.Logger
}

func NewGetDirectMessagesController(
	usecase *message_application.GetDirectMessageHistoryUsecase,
	logger pkg.Logger,
) *GetDirectMessagesController {
	return &GetDirectMessagesController{
		usecase: usecase,
		logger:  logger,
	}
}

// GetDirectMessages godoc
// @Summary Get direct messages
// @Description Get paginated messages from the direct
// @Tags messages
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Param cursor query string false "Cursor Message ID"
// @Param dir query string false "Scrolling direction" default("older")
// @Param count query int false "Number of messages per page" default(10)
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string "Invalid parameters"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 403 {object} map[string]string "Access denied to chat"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/messages/direct/{username} [get]
func (c *GetDirectMessagesController) GetDirectMessages(ctx *gin.Context) {
	userId := ctx.GetString("UserID")
	username := ctx.Param("username")

	cursorMessageId := ctx.Query("cursor")
	direction := ctx.Query("dir")

	count, err := strconv.Atoi(ctx.DefaultQuery("count", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid count parameter"})
		return
	}

	messages, err := c.usecase.Execute(
		userId, username, cursorMessageId, count, direction)
	if err != nil {
		if errors.Is(err, user_domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		ctx.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	ctx.JSON(http.StatusOK, messages)
}
