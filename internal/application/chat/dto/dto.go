package dto

import (
	user_application "main/internal/application/user/dto"
	"time"
)

type ChatItem struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	LastMessage *LastMessage `json:"last_message,omitempty"`
}

type LastMessage struct {
	ID        string                   `json:"id"`
	Content   string                   `json:"content"`
	Timestamp time.Time                `json:"timestamp"`
	IsReplied bool                     `json:"is_replied"`
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

type Typing struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	ChatID   string `json:"chat_id"`
	IsTyping bool   `json:"is_typing"`
}
