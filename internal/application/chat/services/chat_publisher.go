package services

import "main/internal/application/chat/dto"

type ChatEvent string

const (
	ChatEventCreated ChatEvent = "chat_created"
	ChatEventUpdated ChatEvent = "chat_updated"
)

type ChatPublisher interface {
	PublishToUser(userId string, chat *dto.ChatItem, event ChatEvent) error
	PublishToUsers(usersId []string, chat *dto.ChatItem, event ChatEvent) error
}
