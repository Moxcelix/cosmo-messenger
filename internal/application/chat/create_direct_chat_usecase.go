package chat_application

import (
	chat_domain "main/internal/domain/chat"
	user_domain "main/internal/domain/user"
)

type CreateDirectChatUsecase struct {
	userRepo user_domain.UserRepository
	chatRepo chat_domain.ChatRepository
	policy   *chat_domain.ChatPolicy
}

func NewCreateDirectChatUsecase(
	userRepo user_domain.UserRepository,
	chatRepo chat_domain.ChatRepository,
	policy *chat_domain.ChatPolicy) *CreateDirectChatUsecase {
	return &CreateDirectChatUsecase{
		userRepo: userRepo,
		chatRepo: chatRepo,
		policy:   policy,
	}
}

func (uc *CreateDirectChatUsecase) Execute(firstMemberId, secondMemberId string) error {
	firstMember, err := uc.userRepo.GetUserById(firstMemberId)
	if err != nil {
		return err
	}

	if firstMember == nil {
		return user_domain.ErrUserNotFound
	}

	secondMember, err := uc.userRepo.GetUserById(secondMemberId)
	if err != nil {
		return err
	}

	if secondMember == nil {
		return user_domain.ErrUserNotFound
	}

	//directChat, err := uc.chatRepo.GetDirectChat(firstMemberId, secondMemberId)

	return nil
}
