package message_api

import (
	message_domain "main/internal/domain/message"
	"time"
)

type messageView struct {
	ID        string    `json:"id"`
	ChatID    string    `json:"chat_id"`
	SenderID  string    `json:"sender_id"`
	Content   string    `json:"content"`
	ReplyTo   string    `json:"reply_to,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Edited    bool      `json:"edited"`
}

func ToMessageView(msg *message_domain.Message) *messageView {
	time := msg.CreatedAt
	edited := !msg.UpdatedAt.Equal(msg.CreatedAt)

	if edited {
		time = msg.UpdatedAt
	}

	return &messageView{
		ID:        msg.ID,
		ChatID:    msg.ChatID,
		SenderID:  msg.SenderID,
		Content:   msg.Content,
		ReplyTo:   msg.ReplyTo,
		Timestamp: time,
		Edited:    edited,
	}
}

func ToMessageViews(messages []*message_domain.Message) []*messageView {
	result := make([]*messageView, len(messages))
	for i, msg := range messages {
		result[i] = ToMessageView(msg)
	}
	return result
}
