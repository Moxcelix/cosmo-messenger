package message_application

type MessageBroadcaster interface {
	BroadcastToUser(userId string, msg *ChatMessage) error
	BroadcastToUsers(usersId []string, msg *ChatMessage) error
}
