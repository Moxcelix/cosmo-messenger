package factories

import (
	auth_models "main/internal_new/domain/auth/models"
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/projections"
)

type ChatProjectionFactory struct {
}

func NewChatProjectionFactory() *ChatProjectionFactory {
	return &ChatProjectionFactory{}
}

func (s *ChatProjectionFactory) ProjectChat(
	chat *models.Chat,
	users map[string]*auth_models.User,
	requestingUserId string,
) *projections.ChatProjection {
	if chat.Type == models.ChatTypeDirect {
		companionId := chat.GetMemberIdsExcluding(requestingUserId)[0]
		companion := users[companionId]
		requester := users[requestingUserId]

		return s.ProjectDirect(chat, requester, companion)
	}

	return &projections.ChatProjection{
		ID:   chat.ID,
		Name: chat.Name,
		Type: string(chat.Type),
	}
}

func (s *ChatProjectionFactory) ProjectDirect(
	chat *models.Chat,
	requester *auth_models.User,
	companion *auth_models.User,
) *projections.ChatProjection {
	chatName := companion.Name

	return &projections.ChatProjection{
		ID:   chat.ID,
		Name: chatName,
		Type: string(chat.Type),
	}
}

func (s *ChatProjectionFactory) ProjectChats(
	chats map[string]*models.Chat,
	users map[string]*auth_models.User,
	requestingUserId string,
) map[string]*projections.ChatProjection {
	chatHeaders := make(map[string]*projections.ChatProjection)

	for chatId, chat := range chats {
		chatHeader := s.ProjectChat(chat, users, requestingUserId)
		chatHeaders[chatId] = chatHeader
	}

	return chatHeaders
}
