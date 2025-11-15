package broadcasters

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/services"
	"main/pkg"
)

type WebsocketTypingPublisher struct {
	wsHub *pkg.WebSocketHub
}

func NewWebsocketTypingPublisher(
	wsHub *pkg.WebSocketHub,
) services.TypingPublisher {
	return &WebsocketTypingPublisher{
		wsHub: wsHub,
	}
}

func (b *WebsocketTypingPublisher) PublishToUser(
	userId string, typing *dto.Typing) error {
	b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
		Type:    "user_typing",
		Payload: typing,
	})

	return nil
}

func (b *WebsocketTypingPublisher) PublishToUsers(
	usersId []string, typing *dto.Typing) error {
	for _, userId := range usersId {
		b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
			Type:    "user_typing",
			Payload: typing,
		})
	}

	return nil
}
