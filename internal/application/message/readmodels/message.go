package readmodels

import (
	user_readmodels "main/internal/application/user/readmodels"
	"time"
)

type Message struct {
	ID          string
	Content     string
	ChatID      string
	ReplyTo     *Reply
	Sender      *user_readmodels.Sender
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Attachments []*Attachment
}
