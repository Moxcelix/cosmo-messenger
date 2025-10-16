package chat_domain

import "time"

type ChatRepository interface {
	Create(chat *Chat) error
	GetByID(id string) (*Chat, error)
	GetUserChats(userID string, offset, limit int) (*ChatList, error)
	FindUserChat(userID, keyWord string, offset, limit int) (*ChatList, error)
	GetDirectChat(firstUserID, secondUserID string) (*Chat, error)
	Update(chat *Chat) error
	Delete(id string) error
	MarkUpdated(chatID string, updateTime time.Time) error
	ChatExists(chatId string) (bool, error)
	DirectChatExists(firstUserID, secondUserID string) (bool, error)
	UserInChat(userId, chatId string) (bool, error)
}
