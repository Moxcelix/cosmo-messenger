package broadcasters

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/services"
	"main/pkg"
)

type WebsocketChatPublisher struct {
	wsHub *pkg.WebSocketHub
}

func NewWebsocketChatPublisher(
	wsHub *pkg.WebSocketHub,
) services.ChatPublisher {
	return &WebsocketChatPublisher{
		wsHub: wsHub,
	}
}

func (b *WebsocketChatPublisher) PublishToUser(
	userId string, chat *dto.ChatItem, event services.ChatEvent) error {
	b.wsHub.SendToClient(userId, pkg.WebSocketEvent{
		Type:    string(event),
		Payload: chat,
	})

	return nil
}

func (b *WebsocketChatPublisher) PublishToUsers(
	usersId []string, chat *dto.ChatItem, event services.ChatEvent) error {
	for _, userId := range usersId {
		b.PublishToUser(userId, chat, event)
	}

	return nil
}
