package message_infrastructure

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	"main/pkg"
)

type WebsocketBroadcaster struct {
	wsHub *pkg.WebSocketHub
}

func NewWebsocketBroadcaster(
	wsHub *pkg.WebSocketHub,
) services.MessageBroadcaster {
	return &WebsocketBroadcaster{
		wsHub: wsHub,
	}
}

func (b *WebsocketBroadcaster) BroadcastToUser(userId string, msg *dto.ChatMessage) error {
	msgPayload := msg
	b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
		Type:    "new_message",
		Payload: msgPayload,
	})

	return nil
}

func (b *WebsocketBroadcaster) BroadcastToUsers(
	usersId []string, msg *dto.ChatMessage) error {
	msgPayload := msg
	for _, userId := range usersId {
		b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
			Type:    "new_message",
			Payload: msgPayload,
		})
	}

	return nil
}
