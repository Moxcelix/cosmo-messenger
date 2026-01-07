package usecases

import (
	"main/internal_new/domain/chat/factories"
	"main/internal_new/domain/chat/services"
)

type SendDirectMessageUsecase struct {
	messageFactory   *factories.MessageFactory
	directService    *services.DirectChatService
	sendService      *services.SendMessageService
	companionService *services.CompanionService
}

func NewSendDirectMessageUsecase(
	messageFactory *factories.MessageFactory,
	directService *services.DirectChatService,
	sendService *services.SendMessageService,
	companionService *services.CompanionService,
) *SendDirectMessageUsecase {
	return &SendDirectMessageUsecase{
		messageFactory:   messageFactory,
		directService:    directService,
		sendService:      sendService,
		companionService: companionService,
	}
}

func (uc *SendDirectMessageUsecase) Execute(senderId, targetUsername, content string) error {
	companion, err := uc.companionService.GetCompanion(senderId, targetUsername)
	if err != nil {
		return err
	}

	chat, err := uc.directService.GetDirectChat(senderId, companion.ID)
	if err != nil {
		return err
	}

	msg, err := uc.messageFactory.CreateTextMessage(senderId, content)
	if err != nil {
		return err
	}

	if err := uc.sendService.SendMessage(chat, msg); err != nil {
		return err
	}

	return nil
}
