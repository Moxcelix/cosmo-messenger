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

func (uc *FindUserUsecase) Execute(requesterId, targetUsername string) (*UserReview, error) {
	targetUser, err := uc.userRepo.GetUserByUsername(targetUsername)
	if err != nil {
		return nil, err
	}

	if targetUser == nil {
		return nil, user_domain.ErrUserNotFound
	}

	direct, err := uc.chatRepo.GetDirectChat(requesterId, targetUser.ID)
	if err != nil {
		return nil, err
	}

	var directChatId *string
	if direct != nil {
		directChatId = &direct.ID
	}

	return &UserReview{
		Name:         targetUser.Name,
		Username:     targetUser.Username,
		Bio:          targetUser.Bio,
		ID:           targetUser.ID,
		DirectChatId: directChatId,
	}, nil
}
