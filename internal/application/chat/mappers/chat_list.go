package mappers

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/readmodels"
)

type ChatCollectionMapper struct {
	chatItemMapper *ChatItemMapper
}

func NewChatCollectionMapper(chatItemMapper *ChatItemMapper) *ChatCollectionMapper {
	return &ChatCollectionMapper{
		chatItemMapper: chatItemMapper,
	}
}

func (m *ChatCollectionMapper) MapToDTO(
	readModelList *readmodels.ChatList,
	chatNames map[string]string,
) *dto.ChatCollection {

	chatItems := make([]*dto.ChatItem, len(readModelList.Chats))

	for i, chatReadModel := range readModelList.Chats {
		chatItems[i] = m.chatItemMapper.MapToDTO(chatReadModel, chatNames[chatReadModel.ID])
	}

	return &dto.ChatCollection{
		Chats: chatItems,
		Meta: &dto.ScrollingMeta{
			HasPrev: readModelList.HasPrev,
			HasNext: readModelList.HasNext,
		},
	}
}
