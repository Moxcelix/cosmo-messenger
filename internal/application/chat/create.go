package chat_application

import chat_domain "main/internal/domain/chat"

type ChatCreator struct {
	chatRepo        chat_domain.ChatRepository
	chatBroadcaster ChatBroadcaster
	chatAssembler   *ChatItemAssembler
}

func NewChatCreator(
	chatRepo chat_domain.ChatRepository,
	chatBroadcaster ChatBroadcaster,
	chatAssembler *ChatItemAssembler,
) *ChatCreator {
	return &ChatCreator{
		chatRepo:        chatRepo,
		chatBroadcaster: chatBroadcaster,
		chatAssembler:   chatAssembler,
	}
}

func (c *ChatCreator) Create(chat *chat_domain.Chat) error {
	if err := c.chatRepo.Create(chat); err != nil {
		return err
	}

	chatMembersId := chat.GetMembersId()
	for _, chatMemberId := range chatMembersId {
		chatDto, err := c.chatAssembler.Assemble(chat, chatMemberId)
		if err != nil {
			return err
		}
		c.chatBroadcaster.BroadcastToUser(chatMemberId, chatDto, ChatEventCreated)
	}

	return nil
}
