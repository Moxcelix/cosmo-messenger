package chat_infrastructure

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/services"
	"main/pkg"
)

type WebsocketTypingBroadcaster struct {
	wsHub *pkg.WebSocketHub
}

func NewWebsocketTypingBroadcaster(
	wsHub *pkg.WebSocketHub,
) services.TypingBroadcaster {
	return &WebsocketTypingBroadcaster{
		wsHub: wsHub,
	}
}

func (b *WebsocketTypingBroadcaster) BroadcastToUser(
	userId string, typing *dto.Typing) error {
	b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
		Type:    "user_typing",
		Payload: typing,
	})

	return nil
}

func (b *WebsocketTypingBroadcaster) BroadcastToUsers(
	usersId []string, typing *dto.Typing) error {
	for _, userId := range usersId {
		b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
			Type:    "user_typing",
			Payload: typing,
		})
	}

	return nil
}
