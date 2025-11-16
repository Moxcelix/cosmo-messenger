package usecases

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/mappers"
	"main/internal/application/chat/queries"
	"main/internal/application/chat/readmodels"
	chat_domain "main/internal/domain/chat"
)

const (
	defaultPage  = 1
	defaultCount = 10
	maxPageSize  = 100
)

type GetUserChatsUsecase struct {
	chatListQuery queries.ChatListQuery
	chatRepo      chat_domain.ChatRepository
	namingService *chat_domain.ChatNamingService
	mapper        *mappers.ChatCollectionMapper
}

func NewGetUserChatsUsecase(
	chatListQuery queries.ChatListQuery,
	chatRepo chat_domain.ChatRepository,
	namingService *chat_domain.ChatNamingService,
	mapper *mappers.ChatCollectionMapper,
) *GetUserChatsUsecase {
	return &GetUserChatsUsecase{
		chatListQuery: chatListQuery,
		chatRepo:      chatRepo,
		namingService: namingService,
		mapper:        mapper,
	}
}

func (uc *GetUserChatsUsecase) Execute(userID string, page, count int) (*dto.ChatCollection, error) {
	if page < 1 {
		page = defaultPage
	}
	if count < 1 {
		count = defaultCount
	}
	if count > maxPageSize {
		count = maxPageSize
	}
	offset := (page - 1) * count

	readModelList, err := uc.chatListQuery.Query(userID, offset, count)
	if err != nil {
		return nil, err
	}

	chatNames, err := uc.getChatNames(readModelList.Chats, userID)
	if err != nil {
		return nil, err
	}

	return uc.mapper.MapToDTO(readModelList, chatNames), nil
}

func (uc *GetUserChatsUsecase) getChatNames(
	chatReadModels []*readmodels.ChatWithLastMessage,
	currentUserID string,
) (map[string]string, error) {
	domainChats := make([]*chat_domain.Chat, len(chatReadModels))
	for i, chat := range chatReadModels {
		domainChat, err := uc.chatRepo.GetByID(chat.ID)
		if err != nil {
			return nil, err
		}
		domainChats[i] = domainChat
	}

	chatNames := make(map[string]string)
	for _, domainChat := range domainChats {
		displayName, err := uc.namingService.ResolveChatName(domainChat, currentUserID)
		if err != nil {
			return nil, err
		}
		chatNames[domainChat.ID] = displayName
	}

	return chatNames, nil
}
