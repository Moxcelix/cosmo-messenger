package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	user_domain "main/internal/domain/user"
)

type DirectMessageUsecase struct {
	userRepo          user_domain.UserRepository
	directService     *chat_domain.DirectChatService
	sendService       *message_domain.SendMessageService
	messageDispatcher *services.MessageDispatcher
	messagePolicy     *message_domain.MessagePolicy
	messageFactory    *message_domain.MessageFactory
}

func NewDirectMessageUsecase(
	userRepo user_domain.UserRepository,
	directService *chat_domain.DirectChatService,
	sendService *message_domain.SendMessageService,
	messageDispatcher *services.MessageDispatcher,
	messagePolicy *message_domain.MessagePolicy,
	messageFactory *message_domain.MessageFactory,
) *DirectMessageUsecase {
	return &DirectMessageUsecase{
		userRepo:          userRepo,
		directService:     directService,
		sendService:       sendService,
		messageDispatcher: messageDispatcher,
		messageFactory:    messageFactory,
		messagePolicy:     messagePolicy,
	}
}

func (uc *DirectMessageUsecase) Execute(
	senderId, receiverUsername, content string) (*dto.ChatMessage, error) {
	receiver, err := uc.userRepo.GetUserByUsername(receiverUsername)
	if err != nil {
		return nil, err
	}

	if receiver == nil {
		return nil, user_domain.ErrUserNotFound
	}

	chat, err := uc.directService.GetDirectChat(senderId, receiver.ID)
	if err != nil {
		return nil, err
	}

	if err := uc.messagePolicy.ValidateMessageContent(content); err != nil {
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
