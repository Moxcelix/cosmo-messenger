package message_application

import (
	chat_application "main/internal/application/chat"
	chat_domain "main/internal/domain/chat"
	user_domain "main/internal/domain/user"
)

type DirectMessageUsecase struct {
	chatFactory   *chat_domain.ChatFactory
	userRepo      user_domain.UserRepository
	chatRepo      chat_domain.ChatRepository
	messageSender *MessageSender
	chatCreator   *chat_application.ChatCreator
}

func NewDirectMessageUsecase(
	chatFactory *chat_domain.ChatFactory,
	userRepo user_domain.UserRepository,
	chatRepo chat_domain.ChatRepository,
	messageSender *MessageSender,
	chatCreator *chat_application.ChatCreator,
) *DirectMessageUsecase {
	return &DirectMessageUsecase{
		chatFactory:   chatFactory,
		userRepo:      userRepo,
		chatRepo:      chatRepo,
		messageSender: messageSender,
		chatCreator:   chatCreator,
	}
}

func (uc *DirectMessageUsecase) Execute(
	senderId, receiverUsername, content string) (*ChatMessage, error) {
	receiver, err := uc.userRepo.GetUserByUsername(receiverUsername)
	if err != nil {
		return nil, err
	}

	if receiver == nil {
		return nil, user_domain.ErrUserNotFound
	}

	chat, err := uc.findOrCreateDirectChat(senderId, receiver.ID)
	if err != nil {
		return nil, err
	}

	msg, err := uc.messageSender.SendMessageToChat(chat, senderId, content)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (uc *DirectMessageUsecase) findOrCreateDirectChat(
	senderId, receiverId string) (*chat_domain.Chat, error) {
	chat, err := uc.chatRepo.GetDirectChat(senderId, receiverId)
	if err != nil {
		return nil, err
	}
	if chat != nil {
		return chat, nil
	}

	chat, err = uc.chatFactory.CreateDirectChat(senderId, receiverId)
	if err != nil {
		return nil, err
	}

	if err := uc.chatCreator.Create(chat); err != nil {
		return nil, err
	}

	return chat, nil
}
