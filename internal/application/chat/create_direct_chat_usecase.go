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
	if err := uc.checkUserExists(firstMemberId); err != nil {
		return err
	}

	if firstMemberId == secondMemberId {
		return chat_domain.ErrCannotCreateChatWithSelf
	}

	if err := uc.checkUserExists(secondMemberId); err != nil {
		return err
	}

	directChat, err := uc.chatRepo.GetDirectChat(firstMemberId, secondMemberId)
	if err != nil {
		return err
	}
	if directChat != nil {
		return chat_domain.ErrChatAlreadyExsists
	}

	chat := &chat_domain.Chat{
		Type: chat_domain.ChatTypeDirect,
		Members: []*chat_domain.ChatMember{
			{UserID: firstMemberId, Role: chat_domain.RoleMember},
			{UserID: secondMemberId, Role: chat_domain.RoleMember},
		},
	}

	return uc.chatRepo.Create(chat)
}

func (uc *CreateDirectChatUsecase) checkUserExists(userId string) error {
	result, err := uc.userRepo.UserExists(userId)
	if err != nil {
		return err
	}

	if !result {
		return user_domain.ErrUserNotFound
	}

	return nil
}
