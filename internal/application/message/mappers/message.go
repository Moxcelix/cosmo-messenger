package mappers

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/readmodels"

	user_mappers "main/internal/application/user/mappers"
)

type MessageMapper struct {
	replyMapper      *ReplyMapper
	senderMapper     *user_mappers.SenderMapper
	attachmentMapper *AttachmentMapper
}

func NewMessageMapper(
	replyMapper *ReplyMapper,
	senderMapper *user_mappers.SenderMapper,
	attachmentMapper *AttachmentMapper,
) *MessageMapper {
	return &MessageMapper{
		replyMapper:      replyMapper,
		senderMapper:     senderMapper,
		attachmentMapper: attachmentMapper,
	}
}

func (m *MessageMapper) MapToDTO(message *readmodels.Message) *dto.ChatMessage {
	if message == nil {
		return nil
	}

	return &dto.ChatMessage{
		ID:        message.ID,
		Content:   message.Content,
		ChatID:    message.ChatID,
		ReplyTo:   m.replyMapper.MapToDTO(message.ReplyTo),
		Sender:    m.senderMapper.MapToDTO(message.Sender),
		Timestamp: message.CreatedAt,
		Edited:    !message.CreatedAt.Equal(message.UpdatedAt),
		// Attachments: m.attachmentMapper.MapToDTOList(message.Attachments),
	}
}

func (m *MessageMapper) MapToDTOList(messages []*readmodels.Message) []*dto.ChatMessage {
	if messages == nil {
		return nil
	}

	result := make([]*dto.ChatMessage, len(messages))
	for i, message := range messages {
		result[i] = m.MapToDTO(message)
	}
	return result
}
