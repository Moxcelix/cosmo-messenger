package chat_application

import (
	chat_domain "main/internal/domain/chat"
)

type ChatHeaderProvider struct {
	chatRepo       chat_domain.ChatRepository
	namingService  *ChatNamingService
	typingProvider *TypingProvider
}

func NewChatHeaderProvider(
	chatRepo chat_domain.ChatRepository,
	namingService *ChatNamingService,
	typingProvider *TypingProvider,
) *ChatHeaderProvider {
	return &ChatHeaderProvider{
		chatRepo:       chatRepo,
		namingService:  namingService,
		typingProvider: typingProvider,
	}
}

func (p *ChatHeaderProvider) Provide(chatId, currentUserId string) (*ChatHeader, error) {
	chat, err := p.chatRepo.GetByID(chatId)
	if err != nil {
		return nil, err
	}
	chatName, err := p.namingService.ResolveChatName(chat, currentUserId)
	if err != nil {
		return nil, err
	}
	typing, err := p.typingProvider.Provide(chatId)
	if err != nil {
		return nil, err
	}

	return &ChatHeader{
		ID:     chat.ID,
		Name:   chatName,
		Type:   chat.Type,
		Typing: typing,
	}, nil
}
