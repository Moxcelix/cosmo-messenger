package services

import (
	"main/internal_new/domain/chat/errors"
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/repositories"
)

type UserChatService struct {
	chatRepo repositories.ChatRepository
}

func NewUserChatService(
	chatRepo repositories.ChatRepository,
) *UserChatService {
	return &UserChatService{
		chatRepo: chatRepo,
	}
}

func (s *UserChatService) GetChatForUser(chatId, userId string) (*models.Chat, error) {
	chat, err := s.chatRepo.GetChatById(chatId)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		return nil, errors.ErrChatNotFound
	}

	hasAccess := chat.HasMember(userId)
	if !hasAccess {
		return nil, errors.ErrChatAccessDenied
	}

	return chat, nil
}
