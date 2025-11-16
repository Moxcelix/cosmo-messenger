package queries

import "main/internal/application/chat/readmodels"

type ChatQuery interface {
	Query(chatId, userId string) (*readmodels.ChatWithLastMessage, error)
}
