package services

import (
	"main/internal/application/message/dto"
	"main/internal/application/message/mappers"
	"main/internal/application/message/queries"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
)

type MessageSender struct {
	messagePolicy    *message_domain.MessagePolicy
	messageRepo      message_domain.MessageRepository
	messagePublisher MessagePublisher
	messageQuery     queries.MessageQuery
	messageMapper    *mappers.MessageMapper
}

func NewMessageSender(
	messagePolicy *message_domain.MessagePolicy,
	messageRepo message_domain.MessageRepository,
	messagePublisher MessagePublisher,
	messageQuery queries.MessageQuery,
	messageMapper *mappers.MessageMapper,
) *MessageSender {
	return &MessageSender{
		messagePolicy:    messagePolicy,
		messageRepo:      messageRepo,
		messagePublisher: messagePublisher,
		messageQuery:     messageQuery,
		messageMapper:    messageMapper,
	}
}

func (s *MessageSender) SendMessageToChat(
	chat *chat_domain.Chat, senderID, content string) (*dto.ChatMessage, error) {
	if err := s.messagePolicy.ValidateMessageContent(content); err != nil {
		return nil, err
	}

	message := &message_domain.Message{
		ChatID:   chat.ID,
		SenderID: senderID,
		Content:  content,
	}

	if err := s.messageRepo.CreateMessage(message); err != nil {
		return nil, err
	}

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
