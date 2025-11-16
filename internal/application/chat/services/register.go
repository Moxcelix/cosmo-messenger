package services

import (
	"main/internal/application/chat/mappers"
	"main/internal/application/chat/queries"
	chat_domain "main/internal/domain/chat"
)

type ChatRegistryService struct {
	chatRepo          chat_domain.ChatRepository
	chatMapper        *mappers.ChatItemMapper
	chatNamingService *chat_domain.ChatNamingService
	chatPublisher     ChatPublisher
	chatQuery         queries.ChatQuery
}

func NewChatRegistryService(
	chatRepo chat_domain.ChatRepository,
	chatMapper *mappers.ChatItemMapper,
	chatNamingService *chat_domain.ChatNamingService,
	chatBroadcaster ChatPublisher,
	chatQuery queries.ChatQuery,
) *ChatRegistryService {
	return &ChatRegistryService{
		chatRepo:          chatRepo,
		chatPublisher:     chatBroadcaster,
		chatNamingService: chatNamingService,
		chatQuery:         chatQuery,
		chatMapper:        chatMapper,
	}
}

func (c *ChatRegistryService) Register(chat *chat_domain.Chat) error {
	if err := c.chatRepo.Create(chat); err != nil {
		return err
	}

	chatMembersId := chat.GetMembersId()
	for _, chatMemberId := range chatMembersId {

		chatReadmodel, err := c.chatQuery.Query(chat.ID, chatMemberId)
		if err != nil {
			return err
		}

		chatName, err := c.chatNamingService.ResolveChatName(chat, chatMemberId)
		if err != nil {
			return err
		}

		chatDto := c.chatMapper.MapToDTO(chatReadmodel, chatName)

		c.chatPublisher.PublishToUser(chatMemberId, chatDto, ChatEventCreated)
	}

	return nil
}
