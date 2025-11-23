package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	user_domain "main/internal/domain/user"
)

type DirectMessageUsecase struct {
	messageDispatcher *services.MessageDispatcher
	userService       *user_domain.UserService
	chatService       *chat_domain.ChatService
	messagePolicy     *message_domain.MessagePolicy
	messageService    *message_domain.MessageService
	messageFactory    *message_domain.MessageFactory
}

func NewDirectMessageUsecase(
	userService *user_domain.UserService,
	messageDispatcher *services.MessageDispatcher,
	chatService *chat_domain.ChatService,
	messageService *message_domain.MessageService,
	messageFactory *message_domain.MessageFactory,
	messagePolicy *message_domain.MessagePolicy,
) *DirectMessageUsecase {
	return &DirectMessageUsecase{
		userService:       userService,
		messageDispatcher: messageDispatcher,
		chatService:       chatService,
		messageService:    messageService,
		messageFactory:    messageFactory,
		messagePolicy:     messagePolicy,
	}
}

func (uc *DirectMessageUsecase) Execute(
	senderId, receiverUsername, content string) (*dto.ChatMessage, error) {
	receiver, err := uc.userService.GetUserByUsername(receiverUsername)
	if err != nil {
		return nil, err
	}

	chat, err := uc.chatService.GetDirectChat(senderId, receiver.ID)
	if err != nil {
		return nil, err
	}

	if err := uc.messagePolicy.ValidateMessageContent(content); err != nil {
		return nil, err
	}

	msg, err := uc.messageFactory.CreateTextMessage(chat.ID, senderId, content)
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
