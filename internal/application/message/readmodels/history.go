package readmodels

type MessageHistory struct {
	ChatHeader *ChatHeader
	Messages   []*Message
	HasNext    bool
	HasPrev    bool
}
