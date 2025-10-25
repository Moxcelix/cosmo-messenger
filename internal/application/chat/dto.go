package chat_application

import (
	user_application "main/internal/application/user"
	chat_domain "main/internal/domain/chat"
	"time"
)

type ChatItem struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Type        chat_domain.ChatType `json:"type"`
	LastMessage *LastMessage         `json:"last_message,omitempty"`
}

type ChatHeader struct {
	ID   string               `json:"id"`
	Name string               `json:"name"`
	Type chat_domain.ChatType `json:"type"`
}

type LastMessage struct {
	ID        string                   `json:"id"`
	Content   string                   `json:"content"`
	Timestamp time.Time                `json:"timestamp"`
	Sender    *user_application.Sender `json:"sender"`
}

type ScrollingMeta struct {
	HasPrev bool `json:"has_prev"`
	HasNext bool `json:"has_next"`
	Offset  int  `json:"offset"`
	Total   int  `json:"total"`
}

type ChatCollection struct {
	Chats []*ChatItem    `json:"chats"`
	Meta  *ScrollingMeta `json:"meta"`
}
