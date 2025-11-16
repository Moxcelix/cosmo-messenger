package mappers

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/readmodels"
)

type ChatItemMapper struct {
	lastMessageMapper *LastMessageMapper
}

func NewChatItemMapper(lastMessageMapper *LastMessageMapper) *ChatItemMapper {
	return &ChatItemMapper{
		lastMessageMapper: lastMessageMapper,
	}
}

func (m *ChatItemMapper) MapToDTO(
	chatReadModel *readmodels.ChatWithLastMessage,
	chatName string,
) *dto.ChatItem {
	return &dto.ChatItem{
		ID:          chatReadModel.ID,
		Name:        chatName,
		Type:        chatReadModel.Type,
		LastMessage: m.lastMessageMapper.MapToDTO(chatReadModel.LastMessage),
	}
}
