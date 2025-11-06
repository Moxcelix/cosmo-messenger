package message_application

import (
	chat_application "main/internal/application/chat"
	user_application "main/internal/application/user"
	"time"
)

type Reply struct {
	ID      string                   `json:"id"`
	Content string                   `json:"content"`
	Sender  *user_application.Sender `json:"sender"`
}

type ChatMessage struct {
	ID        string                   `json:"id"`
	Content   string                   `json:"content"`
	ChatID    string                   `json:"chat_id"`
	ReplyTo   *Reply                   `json:"reply_to,omitempty"`
	Sender    *user_application.Sender `json:"sender"`
	Timestamp time.Time                `json:"timestamp"`
	Edited    bool                     `json:"edited"`
}

type MessageHistory struct {
	ChatHeader *chat_application.ChatHeader `json:"chat"`
	Messages   []*ChatMessage               `json:"messages"`
	Meta       ScrollingMeta                `json:"meta"`
}

type DirectMessageHistory struct {
	ChatHeader *chat_application.DirectHeader `json:"chat"`
	Messages   []*ChatMessage                 `json:"messages"`
	Meta       ScrollingMeta                  `json:"meta"`
}

type ScrollingMeta struct {
	HasPrev bool `json:"has_prev"`
	HasNext bool `json:"has_next"`
	Offset  int  `json:"offset"`
	Total   int  `json:"total"`
}
