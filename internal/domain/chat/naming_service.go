package chat_domain

import user_domain "main/internal/domain/user"

type ChatNamingService struct {
	userRepo user_domain.UserRepository
}

func NewChatNamingService(userRepo user_domain.UserRepository) *ChatNamingService {
	return &ChatNamingService{
		userRepo: userRepo,
	}
}

func (s *ChatNamingService) ResolveChatName(
	chat *Chat, currentUserID string) (string, error) {

	if chat.Type != "direct" {
		return chat.Name, nil
	}

	otherUserID := getOtherUserID(chat.Members, currentUserID)
	otherUser, err := s.userRepo.GetUserById(otherUserID)
	if err != nil {
		return "", err
	}

	if otherUser != nil {
		return otherUser.Name, nil
	}

	return "UNKNOWN", nil
}

func getOtherUserID(members []*ChatMember, currentUserID string) string {
	for _, member := range members {
		if member.UserID != currentUserID {
			return member.UserID
		}
	}
	return currentUserID
}
