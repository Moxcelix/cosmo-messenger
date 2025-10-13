package chat_domain

type ChatRepository interface {
	Create(chat *Chat) error
	GetByID(id string) (*Chat, error)
	GetByMember(userID string, offset, limit int) ([]*Chat, error)
	Update(chat *Chat) error
	Delete(id string) error
}
