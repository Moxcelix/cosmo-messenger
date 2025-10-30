package message_infrastructure

import (
	message_application "main/internal/application/message"
	"main/pkg"
)

type WebsocketBroadcaster struct {
	wsHub *pkg.WebSocketHub
}

func NewWebsocketBroadcaster(
	wsHub *pkg.WebSocketHub,
) message_application.MessageBroadcaster {
	return &WebsocketBroadcaster{
		wsHub: wsHub,
	}
}

func (b *WebsocketBroadcaster) BroadcastToUser(userId string, msg *message_application.ChatMessage) error {
	msgPayload := msg
	b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
		Type:    "new_message",
		Payload: msgPayload,
	})

	return nil
}

func (b *WebsocketBroadcaster) BroadcastToUsers(
	usersId []string, msg *message_application.ChatMessage) error {
	msgPayload := msg
	for _, userId := range usersId {
		b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
			Type:    "new_message",
			Payload: msgPayload,
		})
	}

	return nil
}
