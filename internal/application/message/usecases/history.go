package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	chat_domain "main/internal/domain/chat"
)

type GetMessageHistoryUsecase struct {
	chatService    *chat_domain.ChatService
	chatPolicy     *chat_domain.ChatPolicy
	historyService *services.MessageHistoryService
}

func NewGetMessageHistoryUsecase(
	chatPolicy *chat_domain.ChatPolicy,
	chatService *chat_domain.ChatService,
	historyService *services.MessageHistoryService,
) *GetMessageHistoryUsecase {
	return &GetMessageHistoryUsecase{
		chatPolicy:     chatPolicy,
		chatService:    chatService,
		historyService: historyService,
	}
}

func (uc *GetMessageHistoryUsecase) Execute(
	userId, chatId, cursorMessageId string, count int, direction string,
) (*dto.MessageHistory, error) {
	chat, err := uc.chatService.GetChatById(chatId)
	if err != nil {
		return nil, err
	}

	if err := uc.chatPolicy.ValidateUserAccess(userId, chat); err != nil {
		return nil, err
	}

	return uc.historyService.GetMessageHistory(userId, cursorMessageId, chat, count, direction)
}
