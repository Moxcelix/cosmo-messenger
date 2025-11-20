package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	"time"
)

type SendMessageUsecase struct {
	chatService       *chat_domain.ChatService
	messageDispatcher *services.MessageDispatcher
	messageService    *message_domain.ChatMessageService
}

func NewSendMessageUsecase(
	chatService *chat_domain.ChatService,
	messageDispatcher *services.MessageDispatcher,
	messageService *message_domain.ChatMessageService,
) *SendMessageUsecase {
	return &SendMessageUsecase{
		chatService:       chatService,
		messageDispatcher: messageDispatcher,
		messageService:    messageService,
	}
}

func (uc *SendMessageUsecase) Execute(senderId, chatId, content string) (*dto.ChatMessage, error) {
	chat, err := uc.chatService.GetChatForUser(chatId, senderId)
	if err != nil {
		return nil, err
	}

	msg, err := uc.messageService.SendMessage(chat, senderId, content, time.Now())
	if err != nil {
		return nil, err
	}

	msgDto, err := uc.messageDispatcher.DispatchMessage(chat, msg)
	if err != nil {
		return nil, err
	}

	return msgDto, nil
}
