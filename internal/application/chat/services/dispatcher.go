package services

import (
	"main/internal/application/chat/mappers"
	"main/internal/application/chat/queries"
	chat_domain "main/internal/domain/chat"
)

type ChatDispatcher struct {
	chatMapper        *mappers.ChatItemMapper
	chatNamingService *ChatNamingService
	chatPublisher     ChatPublisher
	chatQuery         queries.ChatQuery
}

func NewChatDispatcher(
	chatMapper *mappers.ChatItemMapper,
	chatNamingService *ChatNamingService,
	chatBroadcaster ChatPublisher,
	chatQuery queries.ChatQuery,
) *ChatDispatcher {
	return &ChatDispatcher{
		chatPublisher:     chatBroadcaster,
		chatNamingService: chatNamingService,
		chatQuery:         chatQuery,
		chatMapper:        chatMapper,
	}
}

func (c *ChatDispatcher) DispatchChat(chat *chat_domain.Chat) error {

	chatMembersId := chat.GetMembersId()

	chatReadmodel, err := c.chatQuery.Query(chat.ID)
	if err != nil {
		return err
	}

	for _, chatMemberId := range chatMembersId {

		chatName, err := c.chatNamingService.ResolveChatName(chat, chatMemberId)
		if err != nil {
			return err
		}

		chatDto := c.chatMapper.MapToDTO(chatReadmodel, chatName)

		c.chatPublisher.PublishToUser(chatMemberId, chatDto, ChatEventCreated)
	}

	return nil
}
