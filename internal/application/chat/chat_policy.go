package chat_application

import chat_domain "main/internal/domain/chat"

type ChatPolicy struct {
	chatRepo chat_domain.ChatRepository
}

func NewChatPolicy(chatRepo chat_domain.ChatRepository) *ChatPolicy {
	return &ChatPolicy{
		chatRepo: chatRepo,
	}
}

func (p *ChatPolicy) ValidateUserAccess(userId, chatId string) error {
	hasAccess, err := p.chatRepo.UserInChat(userId, chatId)
	if err != nil {
		return err
	}
	if !hasAccess {
		return chat_domain.ErrChatAccessDenied
	}
	return nil
}
