package message_application

import (
	chat_application "main/internal/application/chat"
	message_domain "main/internal/domain/message"
)

const (
	defaultCount = 10
	maxPageSize  = 100
)

type GetChatMessagesUsecase struct {
	msgRepo          message_domain.MessageRepository
	chatPolicy       *chat_application.ChatPolicy
	historyAssembler *MessageHistoryAssembler
}

func NewGetChatMessagesUsecase(
	msgRepo message_domain.MessageRepository,
	chatPolicy *chat_application.ChatPolicy,
	historyAssembler *MessageHistoryAssembler,
) *GetChatMessagesUsecase {
	return &GetChatMessagesUsecase{
		msgRepo:          msgRepo,
		chatPolicy:       chatPolicy,
		historyAssembler: historyAssembler,
	}
}

func (uc *GetChatMessagesUsecase) Execute(
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
