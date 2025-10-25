package chat_domain

import (
	user_domain "main/internal/domain/user"
)

type DirectChatProvider struct {
	userRepo user_domain.UserRepository
	chatRepo ChatRepository
}

func NewDirectChatProvider(
	userRepo user_domain.UserRepository, chatRepo ChatRepository) *DirectChatProvider {
	return &DirectChatProvider{
		userRepo: userRepo,
		chatRepo: chatRepo,
	}
}

func (p *DirectChatProvider) Provide(firstMemberId, secondMemberId string) (*Chat, error) {
	if err := p.checkUserExists(firstMemberId); err != nil {
		return nil, err
	}

	if err := p.checkUserExists(secondMemberId); err != nil {
		return nil, err
	}

	directChat, err := p.chatRepo.GetDirectChat(firstMemberId, secondMemberId)
	if err != nil {
		return nil, err
	}

	if directChat != nil {
		return directChat, nil
	}

	newDirectChat := &Chat{
		Type: ChatTypeDirect,
		Members: []*ChatMember{
			{UserID: firstMemberId, Role: RoleMember},
			{UserID: secondMemberId, Role: RoleMember},
		},
	}

	if err := p.chatRepo.Create(newDirectChat); err != nil {
		return nil, err
	}

	return newDirectChat, nil
}

func (p *DirectChatProvider) checkUserExists(userId string) error {
	exists, err := p.userRepo.UserExists(userId)
	if err != nil {
		return err
	}

	if !exists {
		return user_domain.ErrUserNotFound
	}

	return nil
}
