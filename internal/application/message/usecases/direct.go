package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	user_domain "main/internal/domain/user"
)

type DirectMessageUsecase struct {
	userRepo    user_domain.UserRepository
	chatRepo    chat_domain.ChatRepository
	messageRepo message_domain.MessageRepository
	chatFactory *chat_domain.ChatFactory

	messageDispatcher *services.MessageDispatcher
	messagePolicy     *message_domain.MessagePolicy
	messageFactory    *message_domain.MessageFactory
}

func NewDirectMessageUsecase(
	userRepo user_domain.UserRepository,
	chatRepo chat_domain.ChatRepository,
	messageRepo message_domain.MessageRepository,
	chatFactory *chat_domain.ChatFactory,

	messageDispatcher *services.MessageDispatcher,
	messageFactory *message_domain.MessageFactory,
	messagePolicy *message_domain.MessagePolicy,
) *DirectMessageUsecase {
	return &DirectMessageUsecase{
		userRepo:          userRepo,
		chatRepo:          chatRepo,
		messageRepo:       messageRepo,
		chatFactory:       chatFactory,
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

	chat, err := uc.chatRepo.GetDirectChat(senderId, receiver.ID)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		chat, err = uc.chatFactory.CreateDirectChat(senderId, receiver.ID)
		if err != nil {
			return nil, err
		}
	}

	if err := uc.messagePolicy.ValidateMessageContent(content); err != nil {
		return nil, err
	}

	msg, err := uc.messageFactory.CreateTextMessage(chat.ID, senderId, content)
	if err != nil {
		return nil, err
	}

	if !chat.IsPersisted() {
		if err := uc.chatRepo.Create(chat); err != nil {
			return nil, err
		}
	}

	if err := uc.messageRepo.CreateMessage(msg); err != nil {
		return nil, err
	}

	if err := uc.chatRepo.MarkUpdated(chat.ID, msg.CreatedAt); err != nil {
		return nil, err
	}

	msgDto, err := uc.messageDispatcher.DispatchMessage(chat, msg)
	if err != nil {
		return nil, err
	}

	return msgDto, nil
}
