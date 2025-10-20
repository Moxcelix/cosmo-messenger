package message_application

import "time"

type Sender struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RepliedMessage struct {
	ID      string  `json:"id"`
	Content string  `json:"content"`
	Sender  *Sender `json:"sender"`
}

type ChatMessage struct {
	ID        string          `json:"id"`
	Content   string          `json:"content"`
	ReplyTo   *RepliedMessage `json:"reply_to,omitempty"`
	Sender    *Sender         `json:"sender"`
	Timestamp time.Time       `json:"timestamp"`
	Edited    bool            `json:"edited"`
}

type ChatMessages struct {
	Messages []*ChatMessage `json:"messages"`
	Meta     ScrollingMeta  `json:"meta"`
}

type ScrollingMeta struct {
	HasPrev bool `json:"has_prev"`
	HasNext bool `json:"has_next"`
	Offset  int  `json:"offset"`
	Total   int  `json:"total"`
}
