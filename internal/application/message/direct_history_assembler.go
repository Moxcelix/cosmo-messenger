package message_application

import (
	chat_application "main/internal/application/chat"
	message_domain "main/internal/domain/message"
	user_domain "main/internal/domain/user"
)

type DirectMessageHistoryAssembler struct {
	msgAssembler         *ChatMessageAssembler
	directHeaderProvider *chat_application.DirectHeaderProvider
}

func NewDirectMessageHistoryAssembler(
	msgAssembler *ChatMessageAssembler,
	directHeaderProvider *chat_application.DirectHeaderProvider,
) *DirectMessageHistoryAssembler {
	return &DirectMessageHistoryAssembler{
		msgAssembler:         msgAssembler,
		directHeaderProvider: directHeaderProvider,
	}
}

func (a *DirectMessageHistoryAssembler) Assemble(
	messageList *message_domain.MessageList,
	companion *user_domain.User,
) (*DirectMessageHistory, error) {
	messages := make([]*ChatMessage, 0, len(messageList.Messages))
	for _, msg := range messageList.Messages {
		message, err := a.msgAssembler.Assemble(msg)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	chatHeader, err := a.directHeaderProvider.Provide(companion)
	if err != nil {
		return nil, err
	}
	chatMessages := &DirectMessageHistory{
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

func (a *DirectMessageHistoryAssembler) AssembleEmpty(
	companion *user_domain.User) (*DirectMessageHistory, error) {
	chatHeader, err := a.directHeaderProvider.Provide(companion)
	if err != nil {
		return nil, err
	}
	chatMessages := &DirectMessageHistory{
		ChatHeader: chatHeader,
		Messages:   []*ChatMessage{},
		Meta: ScrollingMeta{
			HasNext: false,
			HasPrev: false,
			Offset:  0,
			Total:   0,
		},
	}
	return chatMessages, nil
}
