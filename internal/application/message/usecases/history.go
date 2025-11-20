package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/services"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
)

type GetMessageHistoryUsecase struct {
	chatService    *chat_domain.ChatService
	historyService *services.MessageHistoryService
}

func NewGetMessageHistoryUsecase(
	msgRepo message_domain.MessageRepository,
	chatService *chat_domain.ChatService,
	historyService *services.MessageHistoryService,
) *GetMessageHistoryUsecase {
	return &GetMessageHistoryUsecase{
		chatService:    chatService,
		historyService: historyService,
	}
}

func (uc *GetMessageHistoryUsecase) Execute(
	userId, chatId, cursorMessageId string, count int, direction string,
) (*dto.MessageHistory, error) {
	chat, err := uc.chatService.GetChatForUser(chatId, userId)
	if err != nil {
		return nil, err
	}

	return uc.historyService.GetMessageHistory(userId, cursorMessageId, chat, count, direction)
}
