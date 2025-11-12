package services

import "main/internal/application/message/dto"

type MessageBroadcaster interface {
	BroadcastToUser(userId string, msg *dto.ChatMessage) error
	BroadcastToUsers(usersId []string, msg *dto.ChatMessage) error
}
