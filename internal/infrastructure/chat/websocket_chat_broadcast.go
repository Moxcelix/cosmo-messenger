package chat_infrastructure

import (
	chat_application "main/internal/application/chat"
	"main/pkg"
)

type WebsocketChatBroadcaster struct {
	wsHub *pkg.WebSocketHub
}

func NewWebsocketChatBroadcaster(
	wsHub *pkg.WebSocketHub,
) chat_application.ChatBroadcaster {
	return &WebsocketChatBroadcaster{
		wsHub: wsHub,
	}
}

func (b *WebsocketChatBroadcaster) BroadcastToUser(
	userId string, chat *chat_application.ChatItem, event chat_application.ChatEvent) error {
	b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
		Type:    string(event),
		Payload: chat,
	})

	return nil
}

func (b *WebsocketChatBroadcaster) BroadcastToUsers(
	usersId []string, chat *chat_application.ChatItem, event chat_application.ChatEvent) error {
	for _, userId := range usersId {
		b.BroadcastToUser(userId, chat, event)
	}

	return nil
}
