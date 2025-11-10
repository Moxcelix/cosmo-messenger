package chat_api

import (
	"main/pkg"
)

type ChatEvents struct {
	typingWebsocket *TypingWebSocket
	logger          pkg.Logger
	wsHub           *pkg.WebSocketHub
}

func NewChatEvents(
	typingWebsocket *TypingWebSocket,
	logger pkg.Logger,
	wsHub *pkg.WebSocketHub,
) *ChatEvents {
	return &ChatEvents{
		typingWebsocket: typingWebsocket,
		logger:          logger,
		wsHub:           wsHub,
	}
}

func (c *ChatEvents) Setup() {
	c.wsHub.On("typing", c.typingWebsocket.Typing)
}
