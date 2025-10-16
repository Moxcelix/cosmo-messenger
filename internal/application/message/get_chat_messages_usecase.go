package message_application

import (
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
)

const (
	defaultPage  = 1
	defaultCount = 10
	maxPageSize  = 100
)

type GetChatMessagesUsecase struct {
	chatRepo chat_domain.ChatRepository
	msgRepo  message_domain.MessageRepository
}

func NewGetChatMessagesUsecase(
	msgRepo message_domain.MessageRepository,
	chatRepo chat_domain.ChatRepository) *GetChatMessagesUsecase {
	return &GetChatMessagesUsecase{
		msgRepo:  msgRepo,
		chatRepo: chatRepo,
	}
}

func (uc *GetChatMessagesUsecase) Execute(
	userId, chatId string, page, count int) (*message_domain.MessageList, error) {
	if err := uc.validateChat(chatId); err != nil {
		return nil, err
	}

	if err := uc.validateUserAccess(userId, chatId); err != nil {
		return nil, err
	}

	if page < 1 {
		page = defaultPage
	}
	if count < 1 {
		count = defaultCount
	}
	if count > maxPageSize {
		count = maxPageSize
	}

	return uc.msgRepo.GetMessagesByChatId(chatId, (page-1)*count, count)
}

func (uc *GetChatMessagesUsecase) validateChat(chatId string) error {
	exists, err := uc.chatRepo.ChatExists(chatId)
	if err != nil {
		return err
	}
	if !exists {
		return chat_domain.ErrChatNotFound
	}
	return nil
}

func (uc *GetChatMessagesUsecase) validateUserAccess(userId, chatId string) error {
	hasAccess, err := uc.chatRepo.UserInChat(userId, chatId)
	if err != nil {
		return err
	}
	if !hasAccess {
		return chat_domain.ErrChatAccessDenied
	}
	return nil
}
