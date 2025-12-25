package queries

import (
	auth_models "main/internal_new/domain/auth/models"
	"main/internal_new/domain/chat/models"
)

type ChatCollectionItemQuery interface {
	Query(
		chatID string,
	) (
		chat *models.Chat,
		lastMessage *models.Message,
		users map[string]*auth_models.User,
		hasNext bool,
		hasPrev bool,
		err error,
	)
}
