package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
)

type SendMessageUsecase struct {
	chatRepo          chat_domain.ChatRepository
	messagePolicy     *message_domain.MessagePolicy
	messageFactory    *message_domain.MessageFactory
	sendService       *message_domain.SendMessageService
	chatPolicy        *chat_domain.ChatPolicy
	messageDispatcher *services.MessageDispatcher
}

func NewSendMessageUsecase(
	chatRepo chat_domain.ChatRepository,
	chatPolicy *chat_domain.ChatPolicy,
	messagePolicy *message_domain.MessagePolicy,
	messageFactory *message_domain.MessageFactory,
	sendService *message_domain.SendMessageService,
	messageDispatcher *services.MessageDispatcher,
) *SendMessageUsecase {
	return &SendMessageUsecase{
		chatRepo:          chatRepo,
		chatPolicy:        chatPolicy,
		messagePolicy:     messagePolicy,
		messageFactory:    messageFactory,
		sendService:       sendService,
		messageDispatcher: messageDispatcher,
	}
}
func (uc *SendMessageUsecase) Execute(senderId, chatId, content string) (*dto.ChatMessage, error) {
	chat, err := uc.chatRepo.GetChatById(chatId)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		return nil, chat_domain.ErrChatNotFound
	}

	if err := uc.messagePolicy.ValidateMessageContent(content); err != nil {
		return nil, err
	}

	if err := uc.chatPolicy.ValidateUserAccess(senderId, chat); err != nil {
		return nil, err
	}

	msg, err := uc.messageFactory.CreateTextMessage(senderId, content)
	if err != nil {
		return nil, err
	}

	if err := uc.sendService.SendMessage(chat, msg); err != nil {
		return nil, err
	}

	msgDto, err := uc.messageDispatcher.DispatchMessage(chat, msg)
	if err != nil {
		return nil, err
	}

	return msgDto, nil
}
