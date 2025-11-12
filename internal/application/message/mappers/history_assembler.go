package mappers

import (
	chat_application "main/internal/application/chat/services"
	"main/internal/application/message/dto"
	chat_domain "main/internal/domain/chat"
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
	chat *chat_domain.Chat,
	currentUserId string,
) (*dto.MessageHistory, error) {
	messages := make([]*dto.ChatMessage, 0, len(messageList.Messages))
	for _, msg := range messageList.Messages {
		message, err := a.msgAssembler.Assemble(msg)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	chatHeader, err := a.chatHeaderProvider.Provide(chat, currentUserId)
	if err != nil {
		return nil, err
	}

	return &dto.MessageHistory{
		ChatHeader: chatHeader,
		Messages:   messages,
		Meta: dto.ScrollingMeta{
			HasNext: messageList.Offset > 0,
			HasPrev: messageList.Offset < messageList.Total-messageList.Limit,
			Offset:  messageList.Offset,
			Total:   messageList.Total,
		},
	}, nil
}
