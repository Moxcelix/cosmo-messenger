package services

import (
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/projections"
)

type MessageDemoService struct {
}

func NewMessageDemoService() *MessageDemoService {
	return &MessageDemoService{}
}

func (s *MessageDemoService) ProjectMessage(msg *models.Message) *projections.MessageDemo {
	return &projections.MessageDemo{
		ID:        msg.ID,
		ChatID:    msg.ChatID,
		Content:   msg.Content,
		IsReply:   msg.ReplyToId != "",
		Timestamp: msg.CreatedAt,
		SenderId:  msg.SenderID,
	}
}

func (s *MessageDemoService) ProjectMessages(messages map[string]*models.Message) map[string]*projections.MessageDemo {
	messageDemos := make(map[string]*projections.MessageDemo)

	for msgId, msg := range messages {
		messageDemo := s.ProjectMessage(msg)
		messageDemos[msgId] = messageDemo
	}

	return messageDemos
}
