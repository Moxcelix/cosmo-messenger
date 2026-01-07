package usecases

import (
	"main/internal_new/domain/chat/factories"
	"main/internal_new/domain/chat/services"
)

type SendChatMessageUsecase struct {
	messageFactory  *factories.MessageFactory
	userChatService *services.UserChatService
	sendService     *services.SendMessageService
}

func NewSendChatMessageUsecase(
	messageFactory *factories.MessageFactory,
	userChatService *services.UserChatService,
	sendService *services.SendMessageService,
) *SendChatMessageUsecase {
	return &SendChatMessageUsecase{
		messageFactory:  messageFactory,
		userChatService: userChatService,
		sendService:     sendService,
	}
}

func (uc *SendChatMessageUsecase) Execute(senderId, chatId, content string) error {
	chat, err := uc.userChatService.GetChatForUser(chatId, senderId)
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
