package message_domain

type MessageFactory struct {
}

func NewMessageFactory() *MessageFactory {
	return &MessageFactory{}
}

func (f *MessageFactory) CreateTextMessage(chatId, senderId, content string) (*Message, error) {
	return &Message{
		ChatID:   chatId,
		SenderID: senderId,
		Content:  content,
	}, nil
}
