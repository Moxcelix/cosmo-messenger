package message_application

import (
	chat_application "main/internal/application/chat"
	message_domain "main/internal/domain/message"
)

const (
	defaultCount = 10
	maxPageSize  = 100
)

type GetMessageHistoryUsecase struct {
	msgRepo          message_domain.MessageRepository
	chatPolicy       *chat_application.ChatPolicy
	historyAssembler *MessageHistoryAssembler
}

func NewGetMessageHistoryUsecase(
	msgRepo message_domain.MessageRepository,
	chatPolicy *chat_application.ChatPolicy,
	historyAssembler *MessageHistoryAssembler,
) *GetMessageHistoryUsecase {
	return &GetMessageHistoryUsecase{
		msgRepo:          msgRepo,
		chatPolicy:       chatPolicy,
		historyAssembler: historyAssembler,
	}
}

func (uc *GetMessageHistoryUsecase) Execute(
	userId, chatId, cursorMessageId string, count int, direction string) (*MessageHistory, error) {

	if err := uc.chatPolicy.ValidateUserAccess(userId, chatId); err != nil {
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

	return uc.historyAssembler.Assemble(messageList)
}
