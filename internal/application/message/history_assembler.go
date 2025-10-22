package message_application

import (
	user_application "main/internal/application/user"
	message_domain "main/internal/domain/message"
)

type MessageHistoryAssembler struct {
	replyProvider  *ReplyProvider
	senderProvider *user_application.SenderProvider
}

func NewMessageHistoryAssembler(
	replyProvider *ReplyProvider,
	senderProvider *user_application.SenderProvider,
) *MessageHistoryAssembler {
	return &MessageHistoryAssembler{
		replyProvider:  replyProvider,
		senderProvider: senderProvider,
	}
}

func (a *MessageHistoryAssembler) Assemble(
	messageList *message_domain.MessageList) (*MessageHistory, error) {
	messages := make([]*ChatMessage, 0, len(messageList.Messages))
	for _, msg := range messageList.Messages {
		sender, err := a.senderProvider.Provide(msg.SenderID)
		if err != nil {
			return nil, err
		}
		repliedMessage, err := a.replyProvider.Provide(msg.ReplyTo)
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
