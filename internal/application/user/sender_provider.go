package user_application

import (
	user_domain "main/internal/domain/user"
)

type SenderProvider struct {
	userRepo user_domain.UserRepository
}

func NewSenderProvider(userRepo user_domain.UserRepository) *SenderProvider {
	return &SenderProvider{
		userRepo: userRepo,
	}
}

func (p *SenderProvider) Provide(userId string) (*Sender, error) {
	user, err := p.userRepo.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return DeletedSender, nil
	}

	return &Sender{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
