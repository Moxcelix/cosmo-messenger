package message_api

import (
	websocket "main/internal/infrastructure/websocket/message"
	"main/pkg"
)

type MessageEvents struct {
	sendMessageWebsocket *websocket.SendMessageWebSocket
	logger               pkg.Logger
	wsHub                *pkg.WebSocketHub
}

func NewMessageEvents(
	sendMessageWebsocket *websocket.SendMessageWebSocket,
	logger pkg.Logger,
	wsHub *pkg.WebSocketHub,
) *MessageEvents {
	return &MessageEvents{
		sendMessageWebsocket: sendMessageWebsocket,
		logger:               logger,
		wsHub:                wsHub,
	}
}

func (c *MessageEvents) Setup() {
	c.wsHub.On("send_message", c.sendMessageWebsocket.SendMessage)
}
