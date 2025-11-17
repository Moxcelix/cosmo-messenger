package queries

import "main/internal/application/message/readmodels"

type HeaderQuery interface {
	Query(chatId string) (*readmodels.ChatHeader, error)
}
