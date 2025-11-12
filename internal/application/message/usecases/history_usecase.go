package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/mappers"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
)

const (
	defaultCount = 10
	maxPageSize  = 100
)

type GetMessageHistoryUsecase struct {
	msgRepo          message_domain.MessageRepository
	chatRepo         chat_domain.ChatRepository
	chatPolicy       *chat_domain.ChatPolicy
	historyAssembler *mappers.MessageHistoryAssembler
}

func NewGetMessageHistoryUsecase(
	msgRepo message_domain.MessageRepository,
	chatRepo chat_domain.ChatRepository,
	chatPolicy *chat_domain.ChatPolicy,
	historyAssembler *mappers.MessageHistoryAssembler,
) *GetMessageHistoryUsecase {
	return &GetMessageHistoryUsecase{
		msgRepo:          msgRepo,
		chatRepo:         chatRepo,
		chatPolicy:       chatPolicy,
		historyAssembler: historyAssembler,
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

	messageList, err := uc.msgRepo.GetMessagesByChatIdScroll(
		chatId, cursorMessageId, count, direction)
	if err != nil {
		return nil, err
	}

	return uc.historyAssembler.Assemble(messageList, chat, userId)
}
