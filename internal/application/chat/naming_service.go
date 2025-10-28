package chat_application

import (
	user_application "main/internal/application/user"
	chat_domain "main/internal/domain/chat"
)

type ChatNamingService struct {
	senderProvider *user_application.SenderProvider
}

func NewChatNamingService(senderProvider *user_application.SenderProvider) *ChatNamingService {
	return &ChatNamingService{
		senderProvider: senderProvider,
	}
}

func (s *ChatNamingService) ResolveChatName(
	chat *chat_domain.Chat, currentUserID string) (string, error) {

	if chat.Type != "direct" {
		return chat.Name, nil
	}

	otherUserID := getOtherUserID(chat.Members, currentUserID)
	otherUser, err := s.senderProvider.Provide(otherUserID)
	if err != nil {
		return "", err
	}

	if otherUser != nil {
		return otherUser.Name, nil
	}

	return "", nil
}

func getOtherUserID(members []*chat_domain.ChatMember, currentUserID string) string {
	for _, member := range members {
		if member.UserID != currentUserID {
			return member.UserID
		}
	}
	return currentUserID
}
