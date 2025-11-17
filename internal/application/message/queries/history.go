package queries

import "main/internal/application/message/readmodels"

type MessageHistoryQuery interface {
	Query(chatId string, cursorMessageId string, count int, direction string) (*readmodels.MessageHistory, error)
}
