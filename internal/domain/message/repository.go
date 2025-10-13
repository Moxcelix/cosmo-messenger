package message_domain

type MessageRepository interface {
	CreateMessage(message *Message) error
	GetMessageById(id string) (*Message, error)
	UpdateMessage(message *Message) error
	DeleteMessage(id string) error
	GetMessagesByChatId(chatId string, offset, limit int) (*MessageList, error)
}
