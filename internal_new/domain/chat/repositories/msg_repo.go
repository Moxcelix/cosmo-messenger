package repositories

import "main/internal_new/domain/chat/models"

type MessageRepository interface {
	CreateMessage(message *models.Message) error
	GetMessageById(id string) (*models.Message, error)
	UpdateMessage(message *models.Message) error
	DeleteMessage(id string) error
	GetLastChatMessage(chatId string) (*models.Message, error)
}
