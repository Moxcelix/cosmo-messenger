package chat_api

import chat_domain "main/internal/domain/chat"

type chatListItemView struct {
	ID   string               `json:"id" bson:"_id"`
	Type chat_domain.ChatType `json:"type" bson:"type"`
	Name string               `json:"name" bson:"name"`
}

func ToChatListItemView(chat *chat_domain.Chat) *chatListItemView {
	return &chatListItemView{
		ID:   chat.ID,
		Type: chat.Type,
		Name: chat.Name,
	}
}

func ToChatListItemViews(chats []*chat_domain.Chat) []*chatListItemView {
	result := make([]*chatListItemView, len(chats))
	for i, chat := range chats {
		result[i] = ToChatListItemView(chat)
	}
	return result
}
