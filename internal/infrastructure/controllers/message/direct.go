package controllers

import (
	message_application "main/internal/application/message/usecases"
	"main/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DirectMessageController struct {
	directMessageUsecase *message_application.DirectMessageUsecase
	logger               pkg.Logger
}

func NewDirectMessageController(
	directMessageUsecase *message_application.DirectMessageUsecase,
	logger pkg.Logger) *DirectMessageController {
	return &DirectMessageController{
		directMessageUsecase: directMessageUsecase,
		logger:               logger,
	}
}

type msgRequest struct {
	ReceiverUsername string `json:"receiver_username" binding:"required"`
	Content          string `json:"content" binding:"required"`
}

// DirectMessage godoc
// @Summary Message sending
// @Description Send direct message
// @Tags messages
// @Accept json
// @Produce json
// @Param input body msgRequest true "User data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Security     BearerAuth
// @Router /api/v1/messages/direct [post]
func (c *DirectMessageController) DirectMessage(ctx *gin.Context) {
	userId := ctx.GetString("UserID")

	var req msgRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg, err := c.directMessageUsecase.Execute(userId, req.ReceiverUsername, req.Content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, msg)
}
