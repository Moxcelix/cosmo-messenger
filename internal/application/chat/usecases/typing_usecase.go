package usecases

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/services"
	chat_domain "main/internal/domain/chat"
	user_domain "main/internal/domain/user"
)

type TypingUsecase struct {
	userRepo        user_domain.UserRepository
	chatRepo        chat_domain.ChatRepository
	typingPublisher services.TypingPublisher
}

func NewTypingUsecase(
	userRepo user_domain.UserRepository,
	chatRepo chat_domain.ChatRepository,
	typingPublisher services.TypingPublisher,
) *TypingUsecase {
	return &TypingUsecase{
		userRepo:        userRepo,
		chatRepo:        chatRepo,
		typingPublisher: typingPublisher,
	}
}

func (uc *TypingUsecase) Execute(userId, chatId string, isTyping bool) error {
	chat, err := uc.chatRepo.GetByID(chatId)
	if err != nil {
		return err
	}

	if chat == nil {
		return chat_domain.ErrChatNotFound
	}

	user, err := uc.userRepo.GetUserById(userId)
	if err != nil {
		return err
	}

	typingDto := &dto.Typing{
		UserID:   user.ID,
		UserName: user.Name,
		ChatID:   chatId,
		IsTyping: isTyping,
	}

	chatMembersId := chat.GetMembersId()
	if err := uc.typingPublisher.PublishToUsers(chatMembersId, typingDto); err != nil {
		return err
	}

	return nil
}
