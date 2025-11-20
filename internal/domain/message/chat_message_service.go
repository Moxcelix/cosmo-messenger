package message_domain

import (
	chat_domain "main/internal/domain/chat"
	"time"
)

type ChatMessageService struct {
	chatRepo      chat_domain.ChatRepository
	messageRepo   MessageRepository
	messagePolicy *MessagePolicy
}

func NewChatMessageService(
	chatRepo chat_domain.ChatRepository,
	messageRepo MessageRepository,
	messagePolicy *MessagePolicy,
) *ChatMessageService {
	return &ChatMessageService{
		chatRepo:      chatRepo,
		messageRepo:   messageRepo,
		messagePolicy: messagePolicy,
	}
}

func (s *ChatMessageService) SendMessage(
	chat *chat_domain.Chat, senderID, content string, sentAt time.Time,
) (*Message, error) {

	if err := s.messagePolicy.ValidateMessageContent(content); err != nil {
		return nil, err
	}

	if !chat.IsPersisted() {
		if err := s.chatRepo.Create(chat); err != nil {
			return nil, err
		}
	}

	msg := &Message{
		ChatID:    chat.ID,
		SenderID:  senderID,
		Content:   content,
		CreatedAt: sentAt,
	}

	if err := s.messageRepo.CreateMessage(msg); err != nil {
		return nil, err
	}

	if err := s.chatRepo.MarkUpdated(chat.ID, sentAt); err != nil {
		return nil, err
	}

	return msg, nil
}
