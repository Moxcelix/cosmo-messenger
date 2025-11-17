package mappers

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/readmodels"
)

type MessageHistoryMapper struct {
	chatHeaderMapper *ChatHeaderMapper
	messageMapper    *MessageMapper
}

func NewMessageHistoryMapper(
	chatHeaderMapper *ChatHeaderMapper,
	messageMapper *MessageMapper,
) *MessageHistoryMapper {
	return &MessageHistoryMapper{
		chatHeaderMapper: chatHeaderMapper,
		messageMapper:    messageMapper,
	}
}

func (m *MessageHistoryMapper) MapToDTO(
	history *readmodels.MessageHistory,
	chatName string,
) *dto.MessageHistory {
	if history == nil {
		return nil
	}

	return &dto.MessageHistory{
		ChatHeader: m.chatHeaderMapper.MapToDTO(history.ChatHeader, chatName),
		Messages:   m.messageMapper.MapToDTOList(history.Messages),
		Meta: dto.ScrollingMeta{
			HasPrev: history.HasPrev,
			HasNext: history.HasNext,
		},
	}
}
