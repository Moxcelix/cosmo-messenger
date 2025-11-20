package services

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/mappers"
	"main/internal/application/message/queries"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
)

type MessageDispatcher struct {
	messagePublisher MessagePublisher
	messageQuery     queries.MessageQuery
	messageMapper    *mappers.MessageMapper
}

func NewMessageDispatcher(
	messagePublisher MessagePublisher,
	messageQuery queries.MessageQuery,
	messageMapper *mappers.MessageMapper,
) *MessageDispatcher {
	return &MessageDispatcher{
		messagePublisher: messagePublisher,
		messageQuery:     messageQuery,
		messageMapper:    messageMapper,
	}
}

func (s *MessageDispatcher) DispatchMessage(
	chat *chat_domain.Chat, message *message_domain.Message) (*dto.ChatMessage, error) {
	msgReadmodel, err := s.messageQuery.Query(message.ID)
	if err != nil {
		return nil, err
	}

	msgDto := s.messageMapper.MapToDTO(msgReadmodel)

	chatMembersId := chat.GetMembersId()
	if err := s.messagePublisher.PublishToUsers(chatMembersId, msgDto); err != nil {
		return nil, err
	}

	return msgDto, nil
}
