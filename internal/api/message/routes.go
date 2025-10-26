package message_api

import (
	auth_api "main/internal/api/auth"
	"main/pkg"
)

type MessageRoutes struct {
	handler                   pkg.RequestHandler
	directMessageController   *DirectMessageController
	getChatMessagesController *GetChatMessagesController
	sendMessageContoller      *SendMessageController
	authMiddleware            *auth_api.AuthMiddleware
}

func NewMessageRoutes(
	handler pkg.RequestHandler,
	directMessageController *DirectMessageController,
	getChatMessagesController *GetChatMessagesController,
	sendMessageContoller *SendMessageController,
	authMiddleware *auth_api.AuthMiddleware,
) *MessageRoutes {
	return &MessageRoutes{
		directMessageController:   directMessageController,
		getChatMessagesController: getChatMessagesController,
		sendMessageContoller:      sendMessageContoller,
		authMiddleware:            authMiddleware,
		handler:                   handler,
	}
}

func (r *MessageRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/messages").
		Use(r.authMiddleware.Handler())

	group.POST("/direct", r.directMessageController.DirectMessage)
	group.POST("/chat/:chat_id", r.sendMessageContoller.SendMessage)
	group.GET("/chat/:chat_id", r.getChatMessagesController.GetChatMessages)
}
