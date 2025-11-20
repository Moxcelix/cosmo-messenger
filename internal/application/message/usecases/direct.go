package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	user_domain "main/internal/domain/user"
	"time"
)

type DirectMessageUsecase struct {
	userRepo          user_domain.UserRepository
	messageDispatcher *services.MessageDispatcher
	chatService       *chat_domain.ChatService
	messageService    *message_domain.ChatMessageService
}

func NewDirectMessageUsecase(
	userRepo user_domain.UserRepository,
	messageDispatcher *services.MessageDispatcher,
	chatService *chat_domain.ChatService,
	messageService *message_domain.ChatMessageService,
) *DirectMessageUsecase {
	return &DirectMessageUsecase{
		userRepo:          userRepo,
		messageDispatcher: messageDispatcher,
		chatService:       chatService,
		messageService:    messageService,
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

	chat, err := uc.chatService.GetDirectChat(senderId, receiver.ID)
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
