package services

import (
	"main/internal_new/domain/chat/projections"
	"main/internal_new/domain/chat/services"
	"main/pkg"
)

type MessageWebsocketPublisher struct {
	wsHub *pkg.WebSocketHub
}

func NewMessageWebsocketPublisher(
	wsHub *pkg.WebSocketHub,
) services.Publisher {
	return &MessageWebsocketPublisher{
		wsHub: wsHub,
	}
}

func (b *MessageWebsocketPublisher) NewMessage(recipientId string, collection *projections.CollectionProjection) error {
	b.wsHub.SendToClient(recipientId, pkg.WebSocketEvent{
		Type:    "new_message",
		Payload: collection,
	})

	return nil
}
