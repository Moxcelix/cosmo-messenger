package controllers

import (
	chat_application "main/internal/application/chat/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TypingController struct {
	typingUsecase *chat_application.TypingUsecase
}

func NewTypingController(typingUsecase *chat_application.TypingUsecase) *TypingController {
	return &TypingController{
		typingUsecase: typingUsecase,
	}
}

type typingRequest struct {
	IsTyping bool   `json:"is_typing"`
	ChatID   string `json:"chat_id"`
}

type typingResponse struct {
	Message string `json:"message"`
}

// Typing godoc
// @Summary User typing
// @Description Represents user typing
// @Tags chats
// @Accept json
// @Produce json
// @Param input body typingRequest true "Typing data"
// @Success 200 {object} typingResponse
// @Failure 400 {object} map[string]string
// @Security     BearerAuth
// @Router /api/v1/chats/typing [post]
func (c *TypingController) Typing(ctx *gin.Context) {
	var req typingRequest

	userId := ctx.GetString("UserID")

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.typingUsecase.Execute(userId, req.ChatID, req.IsTyping); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, typingResponse{
		Message: "typing state updated",
	})
}
