package api

import (
	message_api "main/internal/api/message"
)

type Event interface {
	Setup()
}

type Events []Event

func NewEvents(
	messageEvents *message_api.MessageEvents,
) Events {
	return Events{
		messageEvents,
	}
}

func (e Events) Setup() {
	for _, event := range e {
		event.Setup()
	}
}
