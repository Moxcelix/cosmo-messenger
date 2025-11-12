package usecases

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/mappers"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	user_domain "main/internal/domain/user"
)

type GetDirectMessageHistoryUsecase struct {
	chatFactory      *chat_domain.ChatFactory
	userRepo         user_domain.UserRepository
	msgRepo          message_domain.MessageRepository
	chatRepo         chat_domain.ChatRepository
	chatPolicy       *chat_domain.ChatPolicy
	historyAssembler *mappers.MessageHistoryAssembler
}

func NewGetDirectMessageHistoryUsecase(
	chatFactory *chat_domain.ChatFactory,
	userRepo user_domain.UserRepository,
	msgRepo message_domain.MessageRepository,
	chatRepo chat_domain.ChatRepository,
	chatPolicy *chat_domain.ChatPolicy,
	historyAssembler *mappers.MessageHistoryAssembler,
) *GetDirectMessageHistoryUsecase {
	return &GetDirectMessageHistoryUsecase{
		chatFactory:      chatFactory,
		msgRepo:          msgRepo,
		chatPolicy:       chatPolicy,
		chatRepo:         chatRepo,
		userRepo:         userRepo,
		historyAssembler: historyAssembler,
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
		virtualChat, err := uc.chatFactory.CreateDirectChat(userId, companion.ID)
		if err != nil {
			return nil, err
		}
		chat = virtualChat
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
		chat.ID, cursorMessageId, count, direction)
	if err != nil {
		return nil, err
	}

	return uc.historyAssembler.Assemble(messageList, chat, userId)
}
