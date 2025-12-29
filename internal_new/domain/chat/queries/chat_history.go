package queries

import (
	auth_models "main/internal_new/domain/auth/models"
	"main/internal_new/domain/chat/models"
)

type ChatHistoryQuery interface {
	Query(
		chatId string,
		cursorMessageId string,
		count int,
		direction string,
	) (
		messages map[string]*models.Message,
		replies map[string]*models.Message,
		users map[string]*auth_models.User,
		hasNext bool,
		hasPrev bool,
		err error,
	)
}
