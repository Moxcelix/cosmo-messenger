package message_application

import (
	chat_application "main/internal/application/chat"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	user_domain "main/internal/domain/user"
)

type GetDirectMessageHistoryUsecase struct {
	userRepo         user_domain.UserRepository
	msgRepo          message_domain.MessageRepository
	chatRepo         chat_domain.ChatRepository
	chatPolicy       *chat_application.ChatPolicy
	historyAssembler *DirectMessageHistoryAssembler
}

func NewGetDirectMessageHistoryUsecase(
	userRepo user_domain.UserRepository,
	msgRepo message_domain.MessageRepository,
	chatRepo chat_domain.ChatRepository,
	chatPolicy *chat_application.ChatPolicy,
	historyAssembler *DirectMessageHistoryAssembler,
) *GetDirectMessageHistoryUsecase {
	return &GetDirectMessageHistoryUsecase{
		msgRepo:          msgRepo,
		chatPolicy:       chatPolicy,
		chatRepo:         chatRepo,
		userRepo:         userRepo,
		historyAssembler: historyAssembler,
	}
}

func (uc *GetDirectMessageHistoryUsecase) Execute(
	userId, targetUsername, cursorMessageId string, count int, direction string,
) (*DirectMessageHistory, error) {
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
		return uc.historyAssembler.AssembleEmpty(companion)
	}

	if err := uc.chatPolicy.ValidateUserAccess(userId, chat.ID); err != nil {
		return nil, err
	}

	if count < 1 {
		count = defaultCount
	}
	if count > maxPageSize {
		count = maxPageSize
	}

	messageList, err := uc.msgRepo.GetMessagesByChatIdScroll(
		chat.ID, cursorMessageId, count, direction)
	if err != nil {
		return nil, err
	}

	return uc.historyAssembler.Assemble(messageList, companion)
}
