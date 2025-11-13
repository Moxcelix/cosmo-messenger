package message_api

import (
	controllers "main/internal/infrastructure/controllers/message"
	"main/internal/infrastructure/middlewares"
	"main/pkg"
)

type MessageRoutes struct {
	handler                     pkg.RequestHandler
	directMessageController     *controllers.DirectMessageController
	getChatMessagesController   *controllers.GetChatMessagesController
	getDirectMessagesController *controllers.GetDirectMessagesController
	sendMessageContoller        *controllers.SendMessageController
	authMiddleware              *middlewares.AuthMiddleware
}

func NewMessageRoutes(
	handler pkg.RequestHandler,
	directMessageController *controllers.DirectMessageController,
	getChatMessagesController *controllers.GetChatMessagesController,
	getDirectMessagesController *controllers.GetDirectMessagesController,
	sendMessageContoller *controllers.SendMessageController,
	authMiddleware *middlewares.AuthMiddleware,
) *MessageRoutes {
	return &MessageRoutes{
		directMessageController:     directMessageController,
		getChatMessagesController:   getChatMessagesController,
		sendMessageContoller:        sendMessageContoller,
		getDirectMessagesController: getDirectMessagesController,
		authMiddleware:              authMiddleware,
		handler:                     handler,
	}
}

func (r *MessageRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/messages").
		Use(r.authMiddleware.Handler())

	group.POST("/direct", r.directMessageController.DirectMessage)
	group.POST("/chat/:chat_id", r.sendMessageContoller.SendMessage)

	group.GET("/direct/:username", r.getDirectMessagesController.GetDirectMessages)
	group.GET("/chat/:chat_id", r.getChatMessagesController.GetChatMessages)
}
