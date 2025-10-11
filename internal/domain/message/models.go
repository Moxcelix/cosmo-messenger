package message_domain

import "time"

type MessageStatus string

const (
	MessageStatusSent      MessageStatus = "sent"
	MessageStatusDelivered MessageStatus = "delivered"
	MessageStatusRead      MessageStatus = "read"
)

type Message struct {
	ID        string        `json:"id" bson:"_id"`
	ChatID    string        `json:"chat_id" bson:"chat_id"`
	SenderID  string        `json:"sender_id" bson:"sender_id"`
	Content   string        `json:"content" bson:"content"`
	ReplyTo   string        `json:"reply_to,omitempty" bson:"reply_to,omitempty"`
	Status    MessageStatus `json:"status" bson:"status"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"updated_at"`
}
