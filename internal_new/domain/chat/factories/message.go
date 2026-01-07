package factories

import "main/internal_new/domain/chat/models"

type MessageFactory struct {
}

func NewMessageFactory() *MessageFactory {
	return &MessageFactory{}
}

func (f *MessageFactory) CreateTextMessage(senderId, content string) (*models.Message, error) {
	return &models.Message{
		SenderID: senderId,
		Content:  content,
	}, nil
}
