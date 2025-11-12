package services

import (
	"main/internal/application/chat/dto"
	chat_domain "main/internal/domain/chat"
)

type ChatHeaderProvider struct {
	namingService *ChatNamingService
}

func NewChatHeaderProvider(
	namingService *ChatNamingService,
) *ChatHeaderProvider {
	return &ChatHeaderProvider{
		namingService: namingService,
	}
}

func (p *ChatHeaderProvider) Provide(
	chat *chat_domain.Chat, currentUserId string) (*dto.ChatHeader, error) {
	chatName, err := p.namingService.ResolveChatName(chat, currentUserId)
	if err != nil {
		return nil, err
	}

	return &dto.ChatHeader{
		ID:   chat.ID,
		Name: chatName,
		Type: chat.Type,
	}, nil
}
