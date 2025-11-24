package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	chat_domain "main/internal/domain/chat"
	user_domain "main/internal/domain/user"
)

type GetDirectMessageHistoryUsecase struct {
	userRepo       user_domain.UserRepository
	directService  *chat_domain.DirectChatService
	historyService *services.MessageHistoryService
}

func NewGetDirectMessageHistoryUsecase(
	userRepo user_domain.UserRepository,
	directService *chat_domain.DirectChatService,
	historyService *services.MessageHistoryService,
) *GetDirectMessageHistoryUsecase {
	return &GetDirectMessageHistoryUsecase{
		userRepo:       userRepo,
		directService:  directService,
		historyService: historyService,
	}
}

func (uc *GetDirectMessageHistoryUsecase) Execute(
	userId, targetUsername, cursorMessageId string, count int, direction string,
) (*dto.MessageHistory, error) {
	companion, err := uc.userRepo.GetUserByUsername(targetUsername)
	if err != nil {
		return nil, err
	}

	if companion == nil {
		return nil, user_domain.ErrUserNotFound
	}

	chat, err := uc.directService.GetDirectChat(userId, companion.ID)
	if err != nil {
		return nil, err
	}

	return uc.historyService.GetMessageHistory(userId, cursorMessageId, chat, count, direction)
}
