package message_domain

import "time"

type Attachment struct {
	ID        string    `json:"id" bson:"_id"`
	Type      string    `json:"type" bson:"type"`
	URL       string    `json:"url" bson:"url"`
	Filename  string    `json:"filename,omitempty" bson:"filename,omitempty"`
	Size      int64     `json:"size,omitempty" bson:"size,omitempty"`
	MimeType  string    `json:"mime_type,omitempty" bson:"mime_type,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
}

type Message struct {
	ID          string       `json:"id" bson:"_id"`
	ChatID      string       `json:"chat_id" bson:"chat_id"`
	SenderID    string       `json:"sender_id" bson:"sender_id"`
	Content     string       `json:"content" bson:"content"`
	ReplyTo     string       `json:"reply_to,omitempty" bson:"reply_to,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty" bson:"attachments,omitempty"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type MessageList struct {
	Messages []*Message
	Total    int
	Offset   int
	Limit    int
}
