package chat_application

import (
	chat_domain "main/internal/domain/chat"
	user_domain "main/internal/domain/user"
)

type TypingProvider struct {
	typingService chat_domain.TypingService
	userRepo      user_domain.UserRepository
}

func NewTypingProvider(
	typingService chat_domain.TypingService,
	userRepo user_domain.UserRepository,
) *TypingProvider {
	return &TypingProvider{
		typingService: typingService,
		userRepo:      userRepo,
	}
}

func (p *TypingProvider) Provide(chatId string) ([]*Typing, error) {
	sessions, err := p.typingService.GetActiveSessions(chatId)
	if err != nil {
		return nil, err
	}

	typingList := make([]*Typing, len(sessions))
	for i, session := range sessions {
		user, err := p.userRepo.GetUserById(session.UserID)
		if err != nil {
			return nil, err
		}

		typingList[i] = &Typing{
			UserID:   session.UserID,
			ChatID:   session.ChatID,
			UserName: user.Name,
			IsTyping: true,
		}
	}

	return typingList, nil
}
