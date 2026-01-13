package queries

import (
	"main/internal_new/domain/chat/models"
)

type ChatCollectionQuery interface {
	Query(
		userID string,
		cursorChatID string,
		count int,
		direction string,
	) (
		collection *models.Collection,
		err error,
	)
}
