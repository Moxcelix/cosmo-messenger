package chat_application

import (
	chat_domain "main/internal/domain/chat"
)

type ChatCollectionAssembler struct {
	chatAssembler *ChatItemAssembler
}

func NewChatCollectionAssembler(
	chatAssembler *ChatItemAssembler,
) *ChatCollectionAssembler {
	return &ChatCollectionAssembler{
		chatAssembler: chatAssembler,
	}
}

func (a *ChatCollectionAssembler) Assemble(
	chatList *chat_domain.ChatList, currentUserId string) (*ChatCollection, error) {
	chats := make([]*ChatItem, 0, len(chatList.Chats))
	for _, chat := range chatList.Chats {
		chatItem, err := a.chatAssembler.Assemble(chat, currentUserId)
		if err != nil {
			return nil, err
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
