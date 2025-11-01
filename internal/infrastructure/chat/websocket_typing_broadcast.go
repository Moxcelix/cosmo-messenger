package chat_infrastructure

import (
	chat_application "main/internal/application/chat"
	"main/pkg"
)

type WebsocketTypingBroadcaster struct {
	wsHub *pkg.WebSocketHub
}

func NewWebsocketTypingBroadcaster(
	wsHub *pkg.WebSocketHub,
) chat_application.TypingBroadcaster {
	return &WebsocketTypingBroadcaster{
		wsHub: wsHub,
	}
}

func (b *WebsocketTypingBroadcaster) BroadcastToUser(userId string, isTyping bool) error {
	b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
		Type:    "new_message",
		Payload: isTyping,
	})

	return nil
}

func (b *WebsocketTypingBroadcaster) BroadcastToUsers(
	usersId []string, isTyping bool) error {
	msgPayload := msg
	for _, userId := range usersId {
		b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
			Type:    "new_message",
			Payload: msgPayload,
		})
	}

	return nil
}
