package message_domain

import (
	chat_domain "main/internal/domain/chat"
)

type MessageService struct {
	chatRepo    chat_domain.ChatRepository
	messageRepo MessageRepository
}

func NewMessageService(
	chatRepo chat_domain.ChatRepository,
	messageRepo MessageRepository,
) *MessageService {
	return &MessageService{
		chatRepo:    chatRepo,
		messageRepo: messageRepo,
	}
}

func (s *MessageService) SendMessage(chat *chat_domain.Chat, msg *Message) error {
	if !chat.IsPersisted() {
		if err := s.chatRepo.Create(chat); err != nil {
			return err
		}
	}

	if err := s.messageRepo.CreateMessage(msg); err != nil {
		return err
	}

	if err := s.chatRepo.MarkUpdated(chat.ID, msg.CreatedAt); err != nil {
		return err
	}

	return nil
}
