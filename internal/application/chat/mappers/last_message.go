package mappers

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/readmodels"

	user_mappers "main/internal/application/user/mappers"
)

type LastMessageMapper struct {
	senderMapper *user_mappers.SenderMapper
}

func NewLastMessageMapper(senderMapper *user_mappers.SenderMapper) *LastMessageMapper {
	return &LastMessageMapper{
		senderMapper: senderMapper,
	}
}

func (m *LastMessageMapper) MapToDTO(lastMessage *readmodels.LastMessage) *dto.LastMessage {
	if lastMessage == nil {
		return nil
	}

	return &dto.LastMessage{
		ID:        lastMessage.ID,
		Content:   lastMessage.Content,
		Timestamp: lastMessage.SentAt,
		Sender:    m.senderMapper.MapToDTO(lastMessage.Sender),
	}
}
