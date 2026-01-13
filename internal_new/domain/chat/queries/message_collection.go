package queries

import (
	"main/internal_new/domain/chat/models"
)

type MessageCollectionQuery interface {
	Query(
		chatId string,
		cursorMessageId string,
		count int,
		direction string,
	) (
		collection *models.Collection,
		err error,
	)
}
