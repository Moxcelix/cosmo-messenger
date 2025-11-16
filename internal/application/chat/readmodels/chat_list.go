package readmodels

type ChatList struct {
	Chats   []*ChatWithLastMessage
	HasNext bool
	HasPrev bool
}
