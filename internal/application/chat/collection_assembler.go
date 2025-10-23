package chat_application

import (
	chat_domain "main/internal/domain/chat"
)

type ChatCollectionAssembler struct {
	lastMessageProvider *LastMessageProvider
}

func NewChatCollectionAssembler(
	lastMessageProvider *LastMessageProvider) *ChatCollectionAssembler {
	return &ChatCollectionAssembler{
		lastMessageProvider: lastMessageProvider,
	}
}

func (p *ChatCollectionAssembler) Assemble(
	chatList *chat_domain.ChatList) (*ChatCollection, error) {
	chats := make([]*ChatItem, 0, len(chatList.Chats))
	for _, chat := range chatList.Chats {

		lastMessage, err := p.lastMessageProvider.Provide(chat.ID)
		if err != nil {
			return nil, err
		}

		chatItem := &ChatItem{
			ID:          chat.ID,
			Name:        chat.Name,
			Type:        chat.Type,
			LastMessage: lastMessage,
		}

		chats = append(chats, chatItem)
	}

	return &ChatCollection{
		Chats: chats,
		Meta: &ScrollingMeta{
			HasPrev: chatList.Offset > 0,
			HasNext: chatList.Offset < chatList.Total-chatList.Limit,
			Offset:  chatList.Offset,
			Total:   chatList.Total,
		},
	}, nil
}
