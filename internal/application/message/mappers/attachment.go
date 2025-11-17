package mappers

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/readmodels"
)

type AttachmentMapper struct{}

func NewAttachmentMapper() *AttachmentMapper {
	return &AttachmentMapper{}
}

func (m *AttachmentMapper) MapToDTO(attachment *readmodels.Attachment) *dto.Attachment {
	if attachment == nil {
		return nil
	}

	return &dto.Attachment{
		ID:        attachment.ID,
		Type:      attachment.Type,
		URL:       attachment.URL,
		Filename:  attachment.Filename,
		MimeType:  attachment.MimeType,
		CreatedAt: attachment.CreatedAt,
	}
}

func (m *AttachmentMapper) MapToDTOList(attachments []*readmodels.Attachment) []*dto.Attachment {
	if attachments == nil {
		return nil
	}

	result := make([]*dto.Attachment, len(attachments))
	for i, attachment := range attachments {
		result[i] = m.MapToDTO(attachment)
	}
	return result
}
