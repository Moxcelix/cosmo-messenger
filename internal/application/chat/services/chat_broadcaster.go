package services

import "main/internal/application/chat/dto"

type ChatEvent string

const (
	ChatEventCreated ChatEvent = "chat_created"
	ChatEventUpdated ChatEvent = "chat_updated"
)

type ChatBroadcaster interface {
	BroadcastToUser(userId string, chat *dto.ChatItem, event ChatEvent) error
	BroadcastToUsers(usersId []string, chat *dto.ChatItem, event ChatEvent) error
}
