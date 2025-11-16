package readmodels

import (
	user_readmodels "main/internal/application/user/readmodels"
	"time"
)

type LastMessage struct {
	ID      string
	Content string
	ReplyId string
	SentAt  time.Time
	Sender  *user_readmodels.Sender
}
