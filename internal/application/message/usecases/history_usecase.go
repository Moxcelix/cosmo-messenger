package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/mappers"
	"main/internal/application/message/queries"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
)

const (
	defaultCount = 10
	maxPageSize  = 100
)

type GetMessageHistoryUsecase struct {
	msgRepo           message_domain.MessageRepository
	chatRepo          chat_domain.ChatRepository
	chatPolicy        *chat_domain.ChatPolicy
	historyQuery      queries.MessageHistoryQuery
	historyMapper     *mappers.MessageHistoryMapper
	chatNamingService *chat_domain.ChatNamingService
}

func NewGetMessageHistoryUsecase(
	msgRepo message_domain.MessageRepository,
	chatRepo chat_domain.ChatRepository,
	chatPolicy *chat_domain.ChatPolicy,
	historyQuery queries.MessageHistoryQuery,
	historyMapper *mappers.MessageHistoryMapper,
	chatNamingService *chat_domain.ChatNamingService,
) *GetMessageHistoryUsecase {
	return &GetMessageHistoryUsecase{
		msgRepo:           msgRepo,
		chatRepo:          chatRepo,
		chatPolicy:        chatPolicy,
		historyQuery:      historyQuery,
		historyMapper:     historyMapper,
		chatNamingService: chatNamingService,
	}
}

func (uc *GetMessageHistoryUsecase) Execute(
	userId, chatId, cursorMessageId string, count int, direction string,
) (*dto.MessageHistory, error) {
	chat, err := uc.chatRepo.GetByID(chatId)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		return nil, chat_domain.ErrChatNotFound
	}

	if err := uc.chatPolicy.ValidateUserAccess(userId, chat); err != nil {
		return nil, err
	}

	if count < 1 {
		count = defaultCount
	}
	if count > maxPageSize {
		count = maxPageSize
	}

	historyReadmodel, err := uc.historyQuery.Query(
		chatId, cursorMessageId, count, direction)
	if err != nil {
		return nil, err
	}

	chatName, err := uc.chatNamingService.ResolveChatName(chat, userId)
	if err != nil {
		return nil, err
	}

	historyDto := uc.historyMapper.MapToDTO(historyReadmodel, chatName)

	return historyDto, nil
}
