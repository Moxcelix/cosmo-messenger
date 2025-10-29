package api

import (
	chat_api "main/internal/api/chat"
	message_api "main/internal/api/message"
)

type Event interface {
	Setup()
}

type Events []Event

func NewEvents(
	messageEvents *message_api.MessageEvents,
	chatEvents *chat_api.ChatEvents,
) Events {
	return Events{
		messageEvents,
		chatEvents,
	}
}

func (e Events) Setup() {
	for _, event := range e {
		event.Setup()
	}
}
