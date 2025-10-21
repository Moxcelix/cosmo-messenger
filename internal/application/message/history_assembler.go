package message_application

import (
	message_domain "main/internal/domain/message"
	user_domain "main/internal/domain/user"
)

type MessageHistoryAssembler struct {
	msgRepo  message_domain.MessageRepository
	userRepo user_domain.UserRepository
}

func NewMessageHistoryAssembler(
	msgRepo message_domain.MessageRepository,
	userRepo user_domain.UserRepository) *MessageHistoryAssembler {
	return &MessageHistoryAssembler{
		msgRepo:  msgRepo,
		userRepo: userRepo,
	}
}

func (a *MessageHistoryAssembler) Assemble(
	messageList *message_domain.MessageList) (*MessageHistory, error) {
	messages := make([]*ChatMessage, 0, len(messageList.Messages))
	for _, msg := range messageList.Messages {
		sender, err := a.getSenderModel(msg.SenderID)
		if err != nil {
			return nil, err
		}
		repliedMessage, err := a.getRepliedMessage(msg.ReplyTo)
		if err != nil {
			return nil, err
		}
		timestamp := msg.CreatedAt
		edited := !msg.UpdatedAt.Equal(msg.CreatedAt)

		message := &ChatMessage{
			ID:        msg.ID,
			Content:   msg.Content,
			ReplyTo:   repliedMessage,
			Sender:    sender,
			Timestamp: timestamp,
			Edited:    edited,
		}

		messages = append(messages, message)
	}
	chatMessages := &MessageHistory{
		Messages: messages,
		Meta: ScrollingMeta{
			HasPrev: messageList.Offset > 0,
			HasNext: messageList.Offset < messageList.Total-messageList.Limit,
			Offset:  messageList.Offset,
			Total:   messageList.Total,
		},
	}

	return chatMessages, nil
}

func (a *MessageHistoryAssembler) getRepliedMessage(msgId string) (*Reply, error) {
	message, err := a.msgRepo.GetMessageById(msgId)
	if err == nil {
		return nil, err
	}

	if message == nil {
		return nil, nil
	}

	sender, err := a.getSenderModel(message.SenderID)
	if err == nil {
		return nil, err
	}

	return &Reply{
		ID:      message.ID,
		Content: message.Content,
		Sender:  sender,
	}, nil
}

func (a *MessageHistoryAssembler) getSenderModel(senderId string) (*Sender, error) {
	sender, err := a.userRepo.GetUserById(senderId)
	if err != nil {
		return nil, err
	}

	if sender == nil {
		return nil, nil
	}

	return &Sender{
		ID:   sender.ID,
		Name: sender.Name,
	}, nil
}
