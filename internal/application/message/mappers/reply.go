package mappers

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/readmodels"

	user_mappers "main/internal/application/user/mappers"
)

type ReplyMapper struct {
	senderMapper *user_mappers.SenderMapper
}

func NewReplyMapper(senderMapper *user_mappers.SenderMapper) *ReplyMapper {
	return &ReplyMapper{
		senderMapper: senderMapper,
	}
}

func (m *ReplyMapper) MapToDTO(reply *readmodels.Reply) *dto.Reply {
	if reply == nil {
		return nil
	}

	return &dto.Reply{
		ID:      reply.ID,
		Content: reply.Content,
		Sender:  m.senderMapper.MapToDTO(reply.Sender),
	}
}
