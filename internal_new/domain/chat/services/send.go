package services

import (
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/repositories"
)

type SendMessageService struct {
	chatRepo        repositories.ChatRepository
	msgRepo         repositories.MessageRepository
	msgEnricher     *MessageEnricher
	realTimeService *RealTimeService
}

func NewSendMessageService(
	chatRepo repositories.ChatRepository,
	msgRepo repositories.MessageRepository,
	msgEnricher *MessageEnricher,
	realTimeService *RealTimeService,
) *SendMessageService {
	return &SendMessageService{
		chatRepo:        chatRepo,
		msgRepo:         msgRepo,
		msgEnricher:     msgEnricher,
		realTimeService: realTimeService,
	}
}

func (s *SendMessageService) SendMessage(chat *models.Chat, msg *models.Message) error {
	if !chat.IsPersisted() {
		if err := s.chatRepo.Create(chat); err != nil {
			return err
		}
	}

	if err := msg.BindToChat(chat); err != nil {
		return err
	}

	if err := s.msgRepo.CreateMessage(msg); err != nil {
		return err
	}

	if err := s.chatRepo.MarkUpdated(chat.ID, msg.CreatedAt); err != nil {
		return err
	}

	collection, err := s.msgEnricher.EnrichMessageWithChat(chat, msg)
	if err != nil {
		return err
	}

	recipientsId := chat.GetMemberIdsExcluding(msg.SenderID)
	if err := s.realTimeService.NewMessage(recipientsId, collection); err != nil {
		return err
	}

	return nil
}
