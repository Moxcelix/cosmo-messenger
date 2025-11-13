package chat_api

import (
	websocket "main/internal/infrastructure/websocket/chat"
	"main/pkg"
)

type ChatEvents struct {
	typingWebsocket *websocket.TypingWebSocket
	logger          pkg.Logger
	wsHub           *pkg.WebSocketHub
}

func NewChatEvents(
	typingWebsocket *websocket.TypingWebSocket,
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
