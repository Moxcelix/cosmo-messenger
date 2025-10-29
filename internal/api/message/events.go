package message_api

import (
	"main/pkg"
)

type MessageEvents struct {
	sendMessageWebsocket *SendMessageWebSocket
	logger               pkg.Logger
	wsHub                *pkg.WebSocketHub
}

func NewMessageEvents(
	sendMessageWebsocket *SendMessageWebSocket,
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
