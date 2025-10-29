package message_domain

type MessageDeliver interface {
	Deliver(msg *Message)
}
