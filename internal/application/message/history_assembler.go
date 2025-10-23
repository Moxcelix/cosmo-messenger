package message_application

import (
	message_domain "main/internal/domain/message"
)

type MessageHistoryAssembler struct {
	msgAssembler *ChatMessageAssembler
}

func NewMessageHistoryAssembler(
	msgAssembler *ChatMessageAssembler,
) *MessageHistoryAssembler {
	return &MessageHistoryAssembler{
		msgAssembler: msgAssembler,
	}
}

func (a *MessageHistoryAssembler) Assemble(
	messageList *message_domain.MessageList) (*MessageHistory, error) {
	messages := make([]*ChatMessage, 0, len(messageList.Messages))
	for _, msg := range messageList.Messages {
		message, err := a.msgAssembler.Assemble(msg)
		if err != nil {
			return nil, err
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
