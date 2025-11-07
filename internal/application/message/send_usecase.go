package message_application

import (
	chat_domain "main/internal/domain/chat"
	"time"
)

type SendMessageUsecase struct {
	chatRepo  chat_domain.ChatRepository
	msgSender *MessageSender
}

func NewSendMessageUsecase(
	chatRepo chat_domain.ChatRepository,
	msgSender *MessageSender,
) *SendMessageUsecase {
	return &SendMessageUsecase{
		chatRepo:  chatRepo,
		msgSender: msgSender,
	}
}

func (uc *SendMessageUsecase) Execute(senderId, chatId, content string) (*ChatMessage, error) {
	chat, err := uc.chatRepo.GetByID(chatId)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		return nil, chat_domain.ErrChatNotFound
	}

	msg, err := uc.msgSender.SendMessageToChat(chat, senderId, content)
	if err != nil {
		return nil, err
	}

	if err := uc.chatRepo.MarkUpdated(chatId, time.Now()); err != nil {
		return nil, err
	}

	return msg, nil
}
