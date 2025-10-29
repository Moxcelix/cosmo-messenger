package message_api

import (
	message_application "main/internal/application/message"
	"main/pkg"
)

type SendMessageWebSocket struct {
	logger             pkg.Logger
	sendMessageUsecase *message_application.SendMessageUsecase
}

type wsMessageRequest struct {
	Content string `json:"content"`
	ChatID  string `json:"chat_id"`
}

func (c *SendMessageWebSocket) SendMessage(clientID string, event pkg.WebSocketEvent) error {
	var msg wsMessageRequest

	if payload, ok := event.Payload.(map[string]any); ok {
		if content, exists := payload["content"]; exists {
			msg.Content = content.(string)
		}
		if chatID, exists := payload["chat_id"]; exists {
			msg.ChatID = chatID.(string)
		}
	}

	if err := c.sendMessageUsecase.Execute(clientID, msg.ChatID, msg.Content); err != nil {
		c.logger.Error("Send message usecase error:", err)
		return err
	}

	c.logger.Info("Message sent via WebSocket", "userID", clientID, "chatID", msg.ChatID)
	return nil
}
