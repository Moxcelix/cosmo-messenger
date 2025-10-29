package message_domain

type MessageBroadcaster interface {
	Broadcast(msg *Message) error
}
