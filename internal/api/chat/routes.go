package chat_api

import (
	controllers "main/internal/infrastructure/controllers/chat"
	"main/internal/infrastructure/middlewares"
	"main/pkg"
)

type ChatRoutes struct {
	handler                pkg.RequestHandler
	getUserChatsController *controllers.GetUserChatsController
	typingController       *controllers.TypingController
	authMiddleware         *middlewares.AuthMiddleware
}

func NewChatRoutes(
	handler pkg.RequestHandler,
	getUserChatsController *controllers.GetUserChatsController,
	typingController *controllers.TypingController,
	authMiddleware *middlewares.AuthMiddleware,
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
