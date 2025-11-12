package services

import (
	"main/internal/application/chat/dto"
	user_application "main/internal/application/user"
	message_domain "main/internal/domain/message"
)

type LastMessageProvider struct {
	msgRepo        message_domain.MessageRepository
	senderProvider *user_application.SenderProvider
}

func NewLastMessageProvider(
	msgRepo message_domain.MessageRepository,
	senderProvider *user_application.SenderProvider,
) *LastMessageProvider {
	return &LastMessageProvider{
		msgRepo:        msgRepo,
		senderProvider: senderProvider,
	}
}

func (p *LastMessageProvider) Provide(chatId string) (*dto.LastMessage, error) {
	msg, err := p.msgRepo.GetLastChatMessage(chatId)
	if err != nil {
		return nil, err
	}

	if msg == nil {
		return nil, nil
	}

	sender, err := p.senderProvider.Provide(msg.SenderID)
	if err != nil {
		return nil, err
	}

	timestamp := msg.CreatedAt
	if !msg.UpdatedAt.Equal(msg.CreatedAt) {
		timestamp = msg.UpdatedAt
	}

	return &dto.LastMessage{
		ID:        msg.ID,
		Content:   msg.Content,
		Timestamp: timestamp,
		Sender:    sender,
	}, nil
}
