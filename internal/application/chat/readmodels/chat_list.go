package readmodels

import "time"

type ChatList struct {
	Chats   []*ChatWithLastMessage
	HasNext bool
	HasPrev bool
}

type ChatWithLastMessage struct {
	ID          string
	Type        string
	Name        string
	Description string
	MemberCount int
	LastMessage *LastMessage
	UpdatedAt   time.Time
}

type LastMessage struct {
	ID      string
	Content string
	ReplyId string
	SentAt  time.Time
	Sender  *LastMessageSender
}

type LastMessageSender struct {
	ID       string
	Name     string
	Username string
}
