package websocket

import (
	"fmt"
	message_application "main/internal/application/message/usecases"
	"main/pkg"
)

type SendMessageWebSocket struct {
	logger             pkg.Logger
	sendMessageUsecase *message_application.SendMessageUsecase
}

func NewSendMessageWebSocket(
	logger pkg.Logger,
	sendMessageUsecase *message_application.SendMessageUsecase,
) *SendMessageWebSocket {
	return &SendMessageWebSocket{
		logger:             logger,
		sendMessageUsecase: sendMessageUsecase,
	}
}

type wsMessageRequest struct {
	Content string `json:"content"`
	ChatID  string `json:"chat_id"`
}

func (c *SendMessageWebSocket) SendMessage(clientID string, event pkg.WebSocketEvent) error {
	var req wsMessageRequest

	if err := pkg.ParsePayload(event.Payload, &req); err != nil {
		return fmt.Errorf("failed to parse typing request: %w", err)
	}

	if _, err := c.sendMessageUsecase.Execute(clientID, req.ChatID, req.Content); err != nil {
		c.logger.Error("Send message usecase error:", err)
		return err
	}

	return nil
}
