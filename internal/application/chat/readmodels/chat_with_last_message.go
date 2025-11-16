package readmodels

import "time"

type ChatWithLastMessage struct {
	ID          string
	Type        string
	Name        string
	Description string
	MemberCount int
	LastMessage *LastMessage
	UpdatedAt   time.Time
}
