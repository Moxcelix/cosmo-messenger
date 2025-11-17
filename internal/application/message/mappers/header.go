package mappers

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/readmodels"
)

type ChatHeaderMapper struct{}

func NewChatHeaderMapper() *ChatHeaderMapper {
	return &ChatHeaderMapper{}
}

func (m *ChatHeaderMapper) MapToDTO(chatHeader *readmodels.ChatHeader, chatName string) *dto.ChatHeader {
	if chatHeader == nil {
		return nil
	}

	return &dto.ChatHeader{
		ID:   chatHeader.ID,
		Type: string(chatHeader.Type),
		Name: chatName,
	}
}
