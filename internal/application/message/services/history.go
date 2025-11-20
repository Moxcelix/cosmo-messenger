package services

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/mappers"
	"main/internal/application/message/queries"
	"main/internal/application/message/readmodels"
	chat_domain "main/internal/domain/chat"
)

type MessageHistoryService struct {
	historyQuery      queries.MessageHistoryQuery
	historyMapper     *mappers.MessageHistoryMapper
	chatNamingService *chat_domain.ChatNamingService
}

func NewMessageHistoryService(
	historyQuery queries.MessageHistoryQuery,
	historyMapper *mappers.MessageHistoryMapper,
	chatNamingService *chat_domain.ChatNamingService,
) *MessageHistoryService {
	return &MessageHistoryService{
		historyQuery:      historyQuery,
		historyMapper:     historyMapper,
		chatNamingService: chatNamingService,
	}
}

func (s *MessageHistoryService) GetMessageHistory(
	userId, cursorMessageId string,
	chat *chat_domain.Chat,
	count int, direction string,
) (*dto.MessageHistory, error) {

	var historyReadmodel *readmodels.MessageHistory

	if chat.IsPersisted() {
		count = s.validateAndLimitCount(count)

		dbModel, err := s.historyQuery.Query(
			chat.ID,
			cursorMessageId,
			count,
			direction,
		)
		if err != nil {
			return nil, err
		}

		historyReadmodel = dbModel

	} else {
		historyReadmodel = &readmodels.MessageHistory{
			ChatHeader: &readmodels.ChatHeader{
				ID:   chat.ID,
				Type: string(chat.Type),
				Name: chat.Name,
			},
			Messages: []*readmodels.Message{},
			HasNext:  false,
			HasPrev:  false,
		}
	}

	chatName, err := s.chatNamingService.ResolveChatName(chat, userId)
	if err != nil {
		return nil, err
	}

	return s.historyMapper.MapToDTO(historyReadmodel, chatName), nil
}

func (s *MessageHistoryService) validateAndLimitCount(count int) int {
	const (
		defaultCount = 10
		maxPageSize  = 100
	)

	if count < 1 {
		count = defaultCount
	}
	if count > maxPageSize {
		count = maxPageSize
	}
	return count
}
