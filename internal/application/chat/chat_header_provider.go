package chat_application

import (
	chat_domain "main/internal/domain/chat"
)

type ChatHeaderProvider struct {
	chatRepo      chat_domain.ChatRepository
	namingService *ChatNamingService
}

func NewChatHeaderProvider(
	chatRepo chat_domain.ChatRepository,
	namingService *ChatNamingService,
) *ChatHeaderProvider {
	return &ChatHeaderProvider{
		chatRepo:      chatRepo,
		namingService: namingService,
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

	return &ChatHeader{
		ID:   chat.ID,
		Name: chatName,
		Type: chat.Type,
	}, nil
}
