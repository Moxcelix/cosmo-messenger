package message_domain

import chat_domain "main/internal/domain/chat"

type SendMessageService struct {
	chatRepo chat_domain.ChatRepository
	msgRepo  MessageRepository
}

func NewSendMessageService(
	chatRepo chat_domain.ChatRepository,
	msgRepo MessageRepository,
) *SendMessageService {
	return &SendMessageService{
		chatRepo: chatRepo,
		msgRepo:  msgRepo,
	}
}

func (s *SendMessageService) SendMessage(chat *chat_domain.Chat, msg *Message) error {
	if !chat.IsPersisted() {
		if err := s.chatRepo.Create(chat); err != nil {
			return err
		}
	}

	if err := msg.BindToChat(chat); err != nil {
		return err
	}

	if err := s.msgRepo.CreateMessage(msg); err != nil {
		return err
	}

	if err := s.chatRepo.MarkUpdated(chat.ID, msg.CreatedAt); err != nil {
		return err
	}

	return nil
}
