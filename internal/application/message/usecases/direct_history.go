package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	chat_domain "main/internal/domain/chat"
	user_domain "main/internal/domain/user"
)

type GetDirectMessageHistoryUsecase struct {
	userRepo       user_domain.UserRepository
	chatRepo       chat_domain.ChatRepository
	chatFactory    *chat_domain.ChatFactory
	historyService *services.MessageHistoryService
}

func NewGetDirectMessageHistoryUsecase(
	userRepo user_domain.UserRepository,
	chatRepo chat_domain.ChatRepository,
	historyService *services.MessageHistoryService,

) *GetDirectMessageHistoryUsecase {
	return &GetDirectMessageHistoryUsecase{
		userRepo:       userRepo,
		chatRepo:       chatRepo,
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

	chat, err := uc.chatRepo.GetDirectChat(userId, companion.ID)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		chat, err = uc.chatFactory.CreateDirectChat(userId, companion.ID)
		if err != nil {
			return nil, err
		}
	}

	return uc.historyService.GetMessageHistory(userId, cursorMessageId, chat, count, direction)
}
