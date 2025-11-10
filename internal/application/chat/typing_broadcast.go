package chat_application

type TypingBroadcaster interface {
	BroadcastToUser(userId string, typing *Typing) error
	BroadcastToUsers(usersId []string, typing *Typing) error
}
