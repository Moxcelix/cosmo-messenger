package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
)

type SendMessageUsecase struct {
	chatService       *chat_domain.ChatService
	chatPolicy        *chat_domain.ChatPolicy
	messagePolicy     *message_domain.MessagePolicy
	messageService    *message_domain.MessageService
	messageFactory    *message_domain.MessageFactory
	messageDispatcher *services.MessageDispatcher
}

func NewSendMessageUsecase(
	chatService *chat_domain.ChatService,
	chatPolicy *chat_domain.ChatPolicy,
	messagePolicy *message_domain.MessagePolicy,
	messageService *message_domain.MessageService,
	messageFactory *message_domain.MessageFactory,
	messageDispatcher *services.MessageDispatcher,
) *SendMessageUsecase {
	return &SendMessageUsecase{
		chatService:       chatService,
		chatPolicy:        chatPolicy,
		messagePolicy:     messagePolicy,
		messageService:    messageService,
		messageFactory:    messageFactory,
		messageDispatcher: messageDispatcher,
	}
}

func (uc *SendMessageUsecase) Execute(senderId, chatId, content string) (*dto.ChatMessage, error) {
	chat, err := uc.chatService.GetChatById(chatId)
	if err != nil {
		return nil, err
	}

	if err := uc.messagePolicy.ValidateMessageContent(content); err != nil {
		return nil, err
	}

	if err := uc.chatPolicy.ValidateUserAccess(senderId, chat); err != nil {
		return nil, err
	}

	msg, err := uc.messageFactory.CreateTextMessage(chatId, senderId, content)
	if err != nil {
		return nil, err
	}

	if err := uc.messageService.SendMessage(chat, msg); err != nil {
		return nil, err
	}

	msgDto, err := uc.messageDispatcher.DispatchMessage(chat, msg)
	if err != nil {
		return nil, err
	}

	return msgDto, nil
}
