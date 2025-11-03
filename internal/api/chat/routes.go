package chat_api

import (
	auth_api "main/internal/api/auth"
	"main/pkg"
)

type ChatRoutes struct {
	handler                pkg.RequestHandler
	getUserChatsController *GetUserChatsController
	typingController       *TypingController
	authMiddleware         *auth_api.AuthMiddleware
}

func NewChatRoutes(
	handler pkg.RequestHandler,
	getUserChatsController *GetUserChatsController,
	typingController *TypingController,
	authMiddleware *auth_api.AuthMiddleware,
) *ChatRoutes {
	return &ChatRoutes{
		getUserChatsController: getUserChatsController,
		typingController:       typingController,
		authMiddleware:         authMiddleware,
		handler:                handler,
	}
}

func (r *ChatRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/chats").
		Use(r.authMiddleware.Handler())

	group.GET("/", r.getUserChatsController.GetUserChats)
	group.POST("/typing", r.typingController.Typing)
}
