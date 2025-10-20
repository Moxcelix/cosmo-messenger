package chat_application

import (
	chat_domain "main/internal/domain/chat"
	"time"
)

type ChatListItem struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Type        chat_domain.ChatType `json:"type"`
	LastMessage *LastMessage         `json:"last_message,omitempty"`
}

type LastMessage struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
	Sender    *Sender   `json:"sender"`
}

type Sender struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PaginationMeta struct {
	HasPrev    bool `json:"has_prev"`
	HasNext    bool `json:"has_next"`
	Page       int  `json:"page"`
	TotalPages int  `json:"total_pages"`
	Total      int  `json:"total"`
	Count      int  `json:"count"`
}

type UserChats struct {
	Chats []*ChatListItem `json:"chats"`
	Meta  PaginationMeta  `json:"meta"`
}