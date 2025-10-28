package chat_application

import (
	chat_domain "main/internal/domain/chat"
)

type ChatItemAssembler struct {
	lastMessageProvider *LastMessageProvider
	namingService       *ChatNamingService
}

func NewChatItemAssembler(
	lastMessageProvider *LastMessageProvider,
	namingService *ChatNamingService,
) *ChatItemAssembler {
	return &ChatItemAssembler{
		lastMessageProvider: lastMessageProvider,
		namingService:       namingService,
	}
}

func (p *ChatItemAssembler) Assemble(chat *chat_domain.Chat, currentUserId string) (*ChatItem, error) {
	lastMessage, err := p.lastMessageProvider.Provide(chat.ID)
	if err != nil {
		return nil, err
	}

	chatName, err := p.namingService.ResolveChatName(chat, currentUserId)
	if err != nil {
		return nil, err
	}

	return &ChatItem{
		ID:          chat.ID,
		Name:        chatName,
		Type:        chat.Type,
		LastMessage: lastMessage,
	}, nil
}
