package chat_application

type ChatEvent string

const (
	ChatEventCreated ChatEvent = "chat_created"
	ChatEventUpdated ChatEvent = "chat_updated"
)

type ChatBroadcaster interface {
	BroadcastToUser(userId string, chat *ChatItem, event ChatEvent) error
	BroadcastToUsers(usersId []string, chat *ChatItem, event ChatEvent) error
}
