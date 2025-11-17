package usecases

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/mappers"
	"main/internal/application/chat/queries"
	chat_domain "main/internal/domain/chat"
)

type GetChatUsecase struct {
	chatQuery     queries.ChatQuery
	chatRepo      chat_domain.ChatRepository
	namingService *chat_domain.ChatNamingService
	mapper        *mappers.ChatItemMapper
	chatPolicy    *chat_domain.ChatPolicy
}

func NewGetChatUsecase(
	chatQuery queries.ChatQuery,
	chatRepo chat_domain.ChatRepository,
	namingService *chat_domain.ChatNamingService,
	mapper *mappers.ChatItemMapper,
	chatPolicy *chat_domain.ChatPolicy,
) *GetChatUsecase {
	return &GetChatUsecase{
		chatQuery:     chatQuery,
		chatRepo:      chatRepo,
		namingService: namingService,
		mapper:        mapper,
		chatPolicy:    chatPolicy,
	}
}

func (uc *GetChatUsecase) Execute(userId, chatId string) (*dto.ChatItem, error) {
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

	chatReadmodel, err := uc.chatQuery.Query(chatId)
	if err != nil {
		return nil, err
	}

	chatName, err := uc.namingService.ResolveChatName(chat, userId)
	if err != nil {
		return nil, err
	}

	return uc.mapper.MapToDTO(chatReadmodel, chatName), nil
}
