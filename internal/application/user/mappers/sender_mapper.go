package mappers

import (
	"main/internal/application/user/dto"
	"main/internal/application/user/readmodels"
)

type SenderMapper struct{}

func NewSenderMapper() *SenderMapper {
	return &SenderMapper{}
}

func (m *SenderMapper) MapToDTO(sender *readmodels.Sender) *dto.Sender {
	if sender == nil {
		return nil
	}

	return &dto.Sender{
		ID:   sender.ID,
		Name: sender.Name,
	}
}
