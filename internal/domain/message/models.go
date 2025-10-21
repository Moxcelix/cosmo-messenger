package message_domain

import "time"

type MessageSender struct {
	Name     string
	Username string
}

type RepliedMessage struct {
	Content  string
	SenderID string
	Sender   *MessageSender
}

type Message struct {
	ID      string
	ChatID  string
	Content string

	SenderID  string
	Sender    *MessageSender
	ReplyToID string
	ReplyTo   *RepliedMessage

	CreatedAt time.Time
	UpdatedAt time.Time
}

type MessageList struct {
	Messages []*Message
	Total    int
	Offset   int
	Limit    int
}
