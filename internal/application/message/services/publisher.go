package services

import "main/internal/application/message/dto"

type MessagePublisher interface {
	PublishToUser(userId string, msg *dto.ChatMessage) error
	PublishToUsers(usersId []string, msg *dto.ChatMessage) error
}
