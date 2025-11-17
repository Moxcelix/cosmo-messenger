package dto

import (
	user_application "main/internal/application/user/dto"
	"time"
)

type ChatHeader struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

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
	ChatHeader *ChatHeader    `json:"chat"`
	Messages   []*ChatMessage `json:"messages"`
	Meta       ScrollingMeta  `json:"meta"`
}

type ScrollingMeta struct {
	HasPrev bool `json:"has_prev"`
	HasNext bool `json:"has_next"`
	Offset  int  `json:"offset"`
	Total   int  `json:"total"`
}

type Attachment struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	URL       string    `json:"url"`
	Filename  string    `json:"filename,omitempty"`
	Size      int64     `json:"size,omitempty"`
	MimeType  string    `json:"mime_type,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
