package services

import (
	auth_models "main/internal_new/domain/auth/models"
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/projections"
)

type ChatHeaderService struct {
}

func NewChatHeaderService() *ChatHeaderService {
	return &ChatHeaderService{}
}

func (s *ChatHeaderService) ProjectChat(
	chat *models.Chat,
	users map[string]*auth_models.User,
	requestingUserId string,
) *projections.ChatHeader {
	chatName := chat.Name

	if chat.Type == models.ChatTypeDirect {
		companionId := chat.GetMemberIdsExcluding(requestingUserId)[0]
		companion := users[companionId]
		chatName = companion.Name
	}

	return &projections.ChatHeader{
		ID:   chat.ID,
		Name: chatName,
		Type: string(chat.Type),
	}
}

func (s *ChatHeaderService) ProjectChats(
	chats map[string]*models.Chat,
	users map[string]*auth_models.User,
	requestingUserId string,
) map[string]*projections.ChatHeader {
	chatHeaders := make(map[string]*projections.ChatHeader)

	for chatId, chat := range chats {
		chatHeader := s.ProjectChat(chat, users, requestingUserId)
		chatHeaders[chatId] = chatHeader
	}

	return chatHeaders
}
