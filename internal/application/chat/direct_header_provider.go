package chat_application

import (
	chat_domain "main/internal/domain/chat"
	user_domain "main/internal/domain/user"
)

type DirectHeaderProvider struct {
	namingService *ChatNamingService
}

func NewDirectHeaderProvider(
	namingService *ChatNamingService,
) *DirectHeaderProvider {
	return &DirectHeaderProvider{
		namingService: namingService,
	}
}

func (p *DirectHeaderProvider) Provide(
	companion *user_domain.User) (*DirectHeader, error) {
	chatName, err := p.namingService.ResolveDirectName(companion)
	if err != nil {
		return nil, err
	}

	return &DirectHeader{
		Username: companion.Username,
		Name:     chatName,
		Type:     chat_domain.ChatTypeDirect,
	}, nil
}
