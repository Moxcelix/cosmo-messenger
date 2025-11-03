package infrastructure

import chat_infrastructure "main/internal/infrastructure/chat"

type Worker interface {
	Run()
}

type Workers []Worker

func NewWorkers(typingService *chat_infrastructure.InMemoryTypingService) Workers {
	return Workers{
		typingService,
	}
}

func (u Workers) Run() {
	for _, worker := range u {
		worker.Run()
	}
}
