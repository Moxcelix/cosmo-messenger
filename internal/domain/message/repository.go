package message_domain

type Populate struct {
	Sender        bool
	Replied       bool
	RepliedSender bool
}

var (
	FullPopulate = &Populate{
		Sender:        true,
		Replied:       true,
		RepliedSender: true,
	}
	CleanPopulate = &Populate{
		Sender:        true,
		Replied:       false,
		RepliedSender: false,
	}
	NoPopulate = &Populate{
		Sender:        false,
		Replied:       false,
		RepliedSender: false,
	}
)

type MessageRepository interface {
	CreateMessage(message *Message) error
	UpdateMessage(message *Message) error
	DeleteMessage(id string) error

	GetMessageById(id string, populate *Populate) (*Message, error)
	GetMessagesByChatId(chatId string, cursor string, limit int, direction string, populate *Populate) (*MessageList, error)
	GetLastChatMessage(chatId string, populate *Populate) (*Message, error)
}
