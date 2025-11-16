package queries

import "main/internal/application/chat/readmodels"

type ChatListQuery interface {
	Query(userId string, offset, limit int) (*readmodels.ChatList, error)
}
