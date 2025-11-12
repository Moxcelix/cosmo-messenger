package services

import "main/internal/application/chat/dto"

type TypingBroadcaster interface {
	BroadcastToUser(userId string, typing *dto.Typing) error
	BroadcastToUsers(usersId []string, typing *dto.Typing) error
}
