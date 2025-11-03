package chat_application

import (
	chat_domain "main/internal/domain/chat"
	user_domain "main/internal/domain/user"
)

type TypingUsecase struct {
	userRepo          user_domain.UserRepository
	chatRepo          chat_domain.ChatRepository
	typingService     chat_domain.TypingService
	typingBroadcaster TypingBroadcaster
}

func NewTypingUsecase(
	userRepo user_domain.UserRepository,
	chatRepo chat_domain.ChatRepository,
	typingService chat_domain.TypingService,
	typingBroadcaster TypingBroadcaster,
) *TypingUsecase {
	return &TypingUsecase{
		userRepo:          userRepo,
		chatRepo:          chatRepo,
		typingService:     typingService,
		typingBroadcaster: typingBroadcaster,
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

	if isTyping {
		if err := uc.typingService.StartTyping(userId, chatId); err != nil {
			return err
		}
	} else {
		if err := uc.typingService.StopTyping(userId, chatId); err != nil {
			return err
		}
	}

	typingDto := &Typing{
		UserID:   user.ID,
		UserName: user.Name,
		ChatID:   chatId,
		IsTyping: isTyping,
	}

	chatMembersId := chat.GetMembersId()
	if err := uc.typingBroadcaster.BroadcastToUsers(chatMembersId, typingDto); err != nil {
		return err
	}

	return nil
}
