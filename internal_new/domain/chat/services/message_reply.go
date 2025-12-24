package services

import (
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/projections"
)

type MessageReplyService struct {
}

func NewMessageReplyService() *MessageReplyService {
	return &MessageReplyService{}
}

func (s *MessageReplyService) ProjectMessage(msg *models.Message) *projections.MessageReply {
	return &projections.MessageReply{
		ID:        msg.ID,
		ChatID:    msg.ChatID,
		SenderID:  msg.SenderID,
		Content:   msg.Content,
		Timestamp: msg.CreatedAt,
	}
}

func (s *MessageReplyService) ProjectMessages(messages map[string]*models.Message) map[string]*projections.MessageReply {
	result := make(map[string]*projections.MessageReply, len(messages))
	for id, msg := range messages {
		if msg == nil {
			continue
		}
		result[id] = s.ProjectMessage(msg)
	}

	return result
}
