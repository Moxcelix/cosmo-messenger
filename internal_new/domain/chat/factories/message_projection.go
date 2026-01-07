package factories

import (
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/projections"
)

const (
	DefaultMessage = false
	ReplyOnly      = true
)

type MessageProjectionFactory struct {
}

func NewMessageProjectionFactory() *MessageProjectionFactory {
	return &MessageProjectionFactory{}
}

func (s *MessageProjectionFactory) ProjectMessage(msg *models.Message, replyOnly bool) *projections.MessageProjction {
	return &projections.MessageProjction{
		ID:        msg.ID,
		ChatID:    msg.ChatID,
		SenderID:  msg.SenderID,
		Content:   msg.Content,
		ReplyToId: msg.ReplyToId,
		Timestamp: msg.CreatedAt,
		IsEdited:  !msg.UpdatedAt.IsZero() && msg.UpdatedAt.After(msg.CreatedAt),
		ReplyOnly: replyOnly,
	}
}

func (s *MessageProjectionFactory) ProjectMessages(
	messages map[string]*models.Message, replyOnly bool) map[string]*projections.MessageProjction {

	result := make(map[string]*projections.MessageProjction, len(messages))
	for id, msg := range messages {
		result[id] = s.ProjectMessage(msg, replyOnly)
	}

	return result
}
