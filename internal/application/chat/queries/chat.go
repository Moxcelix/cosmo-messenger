package queries

import "main/internal/application/chat/readmodels"

type ChatQuery interface {
	Query(chatId string) (*readmodels.ChatWithLastMessage, error)
}
