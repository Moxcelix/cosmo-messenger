package chat_api

import (
	"main/pkg"
)

type TypingWebSocket struct {
}

func NewTypingWebSocket() *TypingWebSocket {
	return &TypingWebSocket{}
}

type wsTypingRequest struct {
	IsTyping bool   `json:"is_typing"`
	ChatID   string `json:"chat_id"`
}

func (c *TypingWebSocket) Typing(clientID string, event pkg.WebSocketEvent) error {
	var req wsTypingRequest

	if payload, ok := event.Payload.(map[string]any); ok {
		if isTyping, exists := payload["is_typing"]; exists {
			req.IsTyping = isTyping.(bool)
		}
		if chatID, exists := payload["chat_id"]; exists {
			req.ChatID = chatID.(string)
		}
	}

	return nil
}
