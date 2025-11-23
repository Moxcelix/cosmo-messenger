package usecases

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/mappers"
	"main/internal/application/chat/queries"
	"main/internal/application/chat/services"
	chat_domain "main/internal/domain/chat"
)

type GetChatUsecase struct {
	chatQuery     queries.ChatQuery
	chatService   *chat_domain.ChatService
	chatPolicy    *chat_domain.ChatPolicy
	namingService *services.ChatNamingService
	mapper        *mappers.ChatItemMapper
}

func NewGetChatUsecase(
	chatQuery queries.ChatQuery,
	chatService *chat_domain.ChatService,
	chatPolicy *chat_domain.ChatPolicy,
	mapper *mappers.ChatItemMapper,
	namingService *services.ChatNamingService,
) *GetChatUsecase {
	return &GetChatUsecase{
		chatQuery:     chatQuery,
		chatService:   chatService,
		namingService: namingService,
		mapper:        mapper,
		chatPolicy:    chatPolicy,
	}
}

func (uc *GetChatUsecase) Execute(userId, chatId string) (*dto.ChatItem, error) {
	chat, err := uc.chatService.GetChatById(chatId)
	if err != nil {
		return nil, err
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
