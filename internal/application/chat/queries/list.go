package queries

import "main/internal/application/chat/readmodels"

type ChatListQuery interface {
	Query(userID string, cursorChatID string, count int, direction string) (*readmodels.ChatList, error)
}
