package chat_api

import (
	"fmt"
	chat_application "main/internal/application/chat/usecases"
	"main/pkg"
)

type TypingWebSocket struct {
	typingUsecase *chat_application.TypingUsecase
}

func NewTypingWebSocket(typingUsecase *chat_application.TypingUsecase) *TypingWebSocket {
	return &TypingWebSocket{
		typingUsecase: typingUsecase,
	}
}

type wsTypingRequest struct {
	IsTyping bool   `json:"is_typing"`
	ChatID   string `json:"chat_id"`
}

func (c *TypingWebSocket) Typing(clientID string, event pkg.WebSocketEvent) error {
	var req wsTypingRequest

	if err := pkg.ParsePayload(event.Payload, &req); err != nil {
		return fmt.Errorf("failed to parse typing request: %w", err)
	}

	return c.typingUsecase.Execute(clientID, req.ChatID, req.IsTyping)
}
