package models

import (
	"main/internal_new/domain/chat/errors"
	"time"
)

type Message struct {
	ID          string       `json:"id" bson:"_id"`
	ChatID      string       `json:"chat_id" bson:"chat_id"`
	SenderID    string       `json:"sender_id" bson:"sender_id"`
	Content     string       `json:"content" bson:"content"`
	ReplyToId   string       `json:"reply_to_id,omitempty" bson:"reply_to_id,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty" bson:"attachments,omitempty"`
	CreatedAt   time.Time    `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time    `bson:"updated_at" json:"updated_at"`
}

func (m *Message) BindToChat(chat *Chat) error {
	if m.ChatID != "" {
		return errors.ErrMessageAlreadyBound
	}

	if !chat.IsPersisted() {
		return errors.ErrChatNotFound
	}

	m.ChatID = chat.ID

	return nil
}

func (m *Message) HasReply() bool {
	return m.ReplyToId != ""
}
