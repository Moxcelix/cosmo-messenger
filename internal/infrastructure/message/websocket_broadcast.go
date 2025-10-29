package message_infrastructure

import (
	message_application "main/internal/application/message"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	"main/pkg"
)

type WebsocketBroadcaster struct {
	wsHub        *pkg.WebSocketHub
	chatRepo     chat_domain.ChatRepository
	msgAssembler *message_application.ChatMessageAssembler
}

func NewWebsocketBroadcaster(
	wsHub *pkg.WebSocketHub,
	chatRepo chat_domain.ChatRepository,
	msgAssembler *message_application.ChatMessageAssembler,
) message_domain.MessageBroadcaster {
	return &WebsocketBroadcaster{
		wsHub:        wsHub,
		chatRepo:     chatRepo,
		msgAssembler: msgAssembler,
	}
}

func (b *WebsocketBroadcaster) Broadcast(msg *message_domain.Message) error {
	chat, err := b.chatRepo.GetByID(msg.ChatID)
	if err != nil {
		return err
	}

	if chat == nil {
		return chat_domain.ErrChatNotFound
	}

	msgPayload, err := b.msgAssembler.Assemble(msg)
	if err != nil {
		return err
	}

	for _, member := range chat.Members {
		b.wsHub.SendToClient(member.UserID, pkg.WebSocketEvent{
			Type:    "new_message",
			Payload: msgPayload,
		})
	}

	return nil
}
