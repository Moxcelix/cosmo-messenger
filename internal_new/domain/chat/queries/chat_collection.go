package queries

import (
	auth_models "main/internal_new/domain/auth/models"
	"main/internal_new/domain/chat/models"
)

type ChatCollectionQuery interface {
	Query(
		userID string,
		cursorChatID string,
		count int,
		direction string,
	) (
		chats map[string]*models.Chat,
		lastMessages map[string]*models.Message,
		users map[string]*auth_models.User,
		hasNext bool,
		hasPrev bool,
		err error,
	)
}
