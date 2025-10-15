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
	firstUser, err := uc.userRepo.GetUserById(firstMemberId)
	if err != nil {
		return err
	}

	if firstUser == nil {
		return user_domain.ErrUserNotFound
	}

	if firstMemberId == secondMemberId {
		return chat_domain.ErrCannotCreateChatWithSelf
	}

	secondUser, err := uc.userRepo.GetUserById(secondMemberId)
	if err != nil {
		return err
	}

	if secondUser == nil {
		return user_domain.ErrUserNotFound
	}

	directChat, err := uc.chatRepo.GetDirectChat(firstMemberId, secondMemberId)

	if err != nil {
		return err
	}

	if directChat != nil {
		return chat_domain.ErrChatAlreadyExsists
	}

	firstMember := chat_domain.ChatMember{
		UserID: firstUser.ID,
		Role:   chat_domain.RoleMember,
	}
	secondMember := chat_domain.ChatMember{
		UserID: secondUser.ID,
		Role:   chat_domain.RoleMember,
	}

	chat := &chat_domain.Chat{
		Type:    chat_domain.ChatTypeDirect,
		Members: []chat_domain.ChatMember{firstMember, secondMember},
	}

	if err := uc.chatRepo.Create(chat); err != nil {
		return err
	}

	return nil
}
