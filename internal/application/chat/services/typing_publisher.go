package services

import "main/internal/application/chat/dto"

type TypingPublisher interface {
	PublishToUser(userId string, typing *dto.Typing) error
	PublishToUsers(usersId []string, typing *dto.Typing) error
}
