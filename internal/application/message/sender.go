package message_application

import (
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
)

type MessageSender struct {
	messagePolicy      *message_domain.MessagePolicy
	messageRepo        message_domain.MessageRepository
	messageAssembler   *ChatMessageAssembler
	messageBroadcaster MessageBroadcaster
}

func NewMessageSender(
	messagePolicy *message_domain.MessagePolicy,
	messageRepo message_domain.MessageRepository,
	messageAssembler *ChatMessageAssembler,
	messageBroadcaster MessageBroadcaster,
) *MessageSender {
	return &MessageSender{
		messagePolicy:      messagePolicy,
		messageRepo:        messageRepo,
		messageAssembler:   messageAssembler,
		messageBroadcaster: messageBroadcaster,
	}
}

func (s *MessageSender) SendMessageToChat(
	chat *chat_domain.Chat, senderID, content string) (*ChatMessage, error) {
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

	msgDto, err := s.messageAssembler.Assemble(message)
	if err != nil {
		return nil, err
	}

	chatMembersId := chat.GetMembersId()
	if err := s.messageBroadcaster.BroadcastToUsers(chatMembersId, msgDto); err != nil {
		return nil, err
	}

	return msgDto, nil
}
