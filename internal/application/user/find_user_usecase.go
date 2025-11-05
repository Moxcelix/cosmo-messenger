package user_application

import (
	chat_domain "main/internal/domain/chat"
	user_domain "main/internal/domain/user"
)

type FindUserUsecase struct {
	userRepo user_domain.UserRepository
	chatRepo chat_domain.ChatRepository
}

func NewFindUserUsecase(
	userRepo user_domain.UserRepository,
	chatRepo chat_domain.ChatRepository,
) *FindUserUsecase {
	return &FindUserUsecase{
		userRepo: userRepo,
		chatRepo: chatRepo,
	}
}

func (uc *FindUserUsecase) Execute(requesterUsername, targetUsername string) (*UserReview, error) {
	user, err := uc.userRepo.GetUserByUsername(targetUsername)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, user_domain.ErrUserNotFound
	}

	direct, err := uc.chatRepo.GetDirectChat(requesterUsername, targetUsername)
	if err != nil {
		return nil, err
	}

	var directChatId *string
	if direct != nil {
		directChatId = &direct.ID
	}

	return &UserReview{
		Name:         user.Name,
		Username:     user.Username,
		Bio:          user.Bio,
		ID:           user.ID,
		DirectChatId: directChatId,
	}, nil
}
