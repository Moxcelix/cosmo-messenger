package message_application

import (
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	"time"
)

type SendMessageUsecase struct {
	msgRepo       message_domain.MessageRepository
	chatRepo      chat_domain.ChatRepository
	messagePolicy *message_domain.MessagePolicy
}

func NewSendMessageUsecase(
	msgRepo message_domain.MessageRepository,
	chatRepo chat_domain.ChatRepository,
	messagePolicy *message_domain.MessagePolicy,
) *SendMessageUsecase {
	return &SendMessageUsecase{
		msgRepo:       msgRepo,
		chatRepo:      chatRepo,
		messagePolicy: messagePolicy,
	}
}

func (uc *SendMessageUsecase) Execute(senderId, chatId, content string) error {
	chatExists, err := uc.chatRepo.ChatExists(chatId)
	if err != nil {
		return err
	}

	if !chatExists {
		return chat_domain.ErrChatNotFound
	}

	if err := uc.messagePolicy.ValidateMessageContent(content); err != nil {
		return err
	}

	message := &message_domain.Message{
		ChatID:   chatId,
		SenderID: senderId,
		Content:  content,
	}

	if err := uc.msgRepo.CreateMessage(message); err != nil {
		return err
	}

	if err := uc.chatRepo.MarkUpdated(chatId, time.Now()); err != nil {
		return err
	}

	return nil
}
