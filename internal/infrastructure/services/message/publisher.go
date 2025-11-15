package broadcasters

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	"main/pkg"
)

type MessageWebsocketPublisher struct {
	wsHub *pkg.WebSocketHub
}

func NewMessageWebsocketPublisher(
	wsHub *pkg.WebSocketHub,
) services.MessagePublisher {
	return &MessageWebsocketPublisher{
		wsHub: wsHub,
	}
}

func (b *MessageWebsocketPublisher) PublishToUser(userId string, msg *dto.ChatMessage) error {
	msgPayload := msg
	b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
		Type:    "new_message",
		Payload: msgPayload,
	})

	return nil
}

func (b *MessageWebsocketPublisher) PublishToUsers(
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
