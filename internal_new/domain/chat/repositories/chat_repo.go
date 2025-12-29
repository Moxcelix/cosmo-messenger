package repositories

import (
	"main/internal_new/domain/chat/models"
	"time"
)

type ChatRepository interface {
	Create(chat *models.Chat) error
	GetChatById(id string) (*models.Chat, error)
	GetDirectChat(firstUserID, secondUserID string) (*models.Chat, error)
	Update(chat *models.Chat) error
	Delete(id string) error
	MarkUpdated(chatID string, updateTime time.Time) error
	ChatExists(chatId string) (bool, error)
	DirectChatExists(firstUserID, secondUserID string) (bool, error)
	UserInChat(userId, chatId string) (bool, error)
}
