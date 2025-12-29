package services

import (
	"main/internal_new/domain/chat/factory"
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/repositories"
)

type DirectChatService struct {
	chatRepo    repositories.ChatRepository
	chatFactory *factory.ChatFactory
}

func NewDirectChatService(chatRepo repositories.ChatRepository, chatFactory *factory.ChatFactory) *DirectChatService {
	return &DirectChatService{
		chatRepo:    chatRepo,
		chatFactory: chatFactory,
	}
}

func (s *DirectChatService) GetDirectChat(user1Id, user2Id string) (*models.Chat, error) {
	chat, err := s.chatRepo.GetDirectChat(user1Id, user2Id)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		chat, err = s.chatFactory.CreateDirectChat(user1Id, user2Id)
		if err != nil {
			return nil, err
		}
	}

	return chat, err
}
