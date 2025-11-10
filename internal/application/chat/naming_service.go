package chat_application

import (
	chat_domain "main/internal/domain/chat"
	user_domain "main/internal/domain/user"
)

type ChatNamingService struct {
	userRepo user_domain.UserRepository
}

func NewChatNamingService(userRepo user_domain.UserRepository) *ChatNamingService {
	return &ChatNamingService{
		userRepo: userRepo,
	}
}

func (s *ChatNamingService) ResolveDirectName(companion *user_domain.User) (string, error) {
	if companion != nil {
		return companion.Name, nil
	}

	return "DELETED", nil
}

func (s *ChatNamingService) ResolveChatName(
	chat *chat_domain.Chat, currentUserID string) (string, error) {

	if chat.Type != "direct" {
		return chat.Name, nil
	}

	otherUserID := getOtherUserID(chat.Members, currentUserID)
	otherUser, err := s.userRepo.GetUserById(otherUserID)
	if err != nil {
		return "", err
	}

	return s.ResolveDirectName(otherUser)
}

func getOtherUserID(members []*chat_domain.ChatMember, currentUserID string) string {
	for _, member := range members {
		if member.UserID != currentUserID {
			return member.UserID
		}
	}
	return currentUserID
}
