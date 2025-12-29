package queries

import (
	auth_models "main/internal_new/domain/auth/models"
	"main/internal_new/domain/chat/models"
)

type DirectQuery interface {
	Query(userId string, companionUsername string) (chat *models.Chat, users map[string]*auth_models.User, err error)
}
