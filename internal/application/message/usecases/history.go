package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	chat_domain "main/internal/domain/chat"
)

type GetMessageHistoryUsecase struct {
	chatRepo       chat_domain.ChatRepository
	chatPolicy     *chat_domain.ChatPolicy
	historyService *services.MessageHistoryService
}

func NewGetMessageHistoryUsecase(
	chatRepo chat_domain.ChatRepository,
	chatPolicy *chat_domain.ChatPolicy,
	historyService *services.MessageHistoryService,
) *GetMessageHistoryUsecase {
	return &GetMessageHistoryUsecase{
		chatPolicy:     chatPolicy,
		chatRepo:       chatRepo,
		historyService: historyService,
	}
}

func (uc *GetMessageHistoryUsecase) Execute(
	userId, chatId, cursorMessageId string, count int, direction string,
) (*dto.MessageHistory, error) {
	chat, err := uc.chatRepo.GetChatById(chatId)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		return nil, chat_domain.ErrChatNotFound
	}

	if err := uc.chatPolicy.ValidateUserAccess(userId, chat); err != nil {
		return nil, err
	}

	return uc.historyService.GetMessageHistory(userId, cursorMessageId, chat, count, direction)
}
