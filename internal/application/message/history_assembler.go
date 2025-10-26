package message_application

import (
	chat_application "main/internal/application/chat"
	message_domain "main/internal/domain/message"
)

type MessageHistoryAssembler struct {
	msgAssembler       *ChatMessageAssembler
	chatHeaderProvider *chat_application.ChatHeaderProvider
}

func NewMessageHistoryAssembler(
	msgAssembler *ChatMessageAssembler,
	chatHeaderProvider *chat_application.ChatHeaderProvider,
) *MessageHistoryAssembler {
	return &MessageHistoryAssembler{
		msgAssembler:       msgAssembler,
		chatHeaderProvider: chatHeaderProvider,
	}
}

func (a *MessageHistoryAssembler) Assemble(
	messageList *message_domain.MessageList,
	chatId, currentUserId string,
) (*MessageHistory, error) {
	messages := make([]*ChatMessage, 0, len(messageList.Messages))
	for _, msg := range messageList.Messages {
		message, err := a.msgAssembler.Assemble(msg)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	chatHeader, err := a.chatHeaderProvider.Provide(chatId, currentUserId)
	if err != nil {
		return nil, err
	}
	chatMessages := &MessageHistory{
		ChatHeader: chatHeader,
		Messages:   messages,
		Meta: ScrollingMeta{
			HasNext: messageList.Offset > 0,
			HasPrev: messageList.Offset < messageList.Total-messageList.Limit,
			Offset:  messageList.Offset,
			Total:   messageList.Total,
		},
	}
	return chatMessages, nil
}
