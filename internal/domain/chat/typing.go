package chat_domain

type Typing interface {
	Set(userId, chatId string, state bool)
}
