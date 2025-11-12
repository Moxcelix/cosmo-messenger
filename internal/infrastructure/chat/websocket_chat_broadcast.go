package chat_infrastructure

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/services"
	"main/pkg"
)

type WebsocketChatBroadcaster struct {
	wsHub *pkg.WebSocketHub
}

func NewWebsocketChatBroadcaster(
	wsHub *pkg.WebSocketHub,
) services.ChatBroadcaster {
	return &WebsocketChatBroadcaster{
		wsHub: wsHub,
	}
}

func (b *WebsocketChatBroadcaster) BroadcastToUser(
	userId string, chat *dto.ChatItem, event services.ChatEvent) error {
	b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
		Type:    string(event),
		Payload: chat,
	})

	return nil
}

func (b *WebsocketChatBroadcaster) BroadcastToUsers(
	usersId []string, chat *dto.ChatItem, event services.ChatEvent) error {
	for _, userId := range usersId {
		b.BroadcastToUser(userId, chat, event)
	}

	return nil
}
