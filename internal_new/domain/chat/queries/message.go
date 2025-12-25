package queries

import (
	auth_models "main/internal_new/domain/auth/models"
	"main/internal_new/domain/chat/models"
)

type MessageQuery interface {
	Query(msgId string) (messages map[string]*models.Message, users map[string]*auth_models.User)
}
