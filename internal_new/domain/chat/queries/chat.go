package queries

import (
	auth_models "main/internal_new/domain/auth/models"
	"main/internal_new/domain/chat/models"
)

type ChatQuery interface {
	Query(chatId string) (chat *models.Chat, users map[string]*auth_models.User, err error)
}
