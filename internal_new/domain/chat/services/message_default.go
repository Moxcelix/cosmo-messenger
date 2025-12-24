package services

import (
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/projections"
)

type MessageDefaultService struct {
}

func NewMessageDefaultService() *MessageDefaultService {
	return &MessageDefaultService{}
}

func (s *MessageDefaultService) ProjectMessage(msg *models.Message) *projections.MessageDefault {
	return &projections.MessageDefault{
		ID:        msg.ID,
		ChatID:    msg.ChatID,
		SenderID:  msg.SenderID,
		Content:   msg.Content,
		ReplyToId: msg.ReplyToId,
		Timestamp: msg.CreatedAt,
		IsEdited:  !msg.UpdatedAt.IsZero() && msg.UpdatedAt.After(msg.CreatedAt),
	}
}

func (s *MessageDefaultService) ProjectMessages(messages map[string]*models.Message) map[string]*projections.MessageDefault {
	if len(messages) == 0 {
		return make(map[string]*projections.MessageDefault)
	}

	result := make(map[string]*projections.MessageDefault, len(messages))
	for id, msg := range messages {
		if msg == nil {
			continue
		}
		result[id] = s.ProjectMessage(msg)
	}

	return result
}
