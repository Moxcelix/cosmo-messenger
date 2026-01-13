package services

import (
	auth_repositories "main/internal_new/domain/auth/repositories"
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/repositories"
)

type MessageEnricher struct {
	msgRepo  repositories.MessageRepository
	userRepo auth_repositories.UserRepository
}

func NewMessageEnricher(
	msgRepo repositories.MessageRepository,
	userRepo auth_repositories.UserRepository,
) *MessageEnricher {
	return &MessageEnricher{
		msgRepo:  msgRepo,
		userRepo: userRepo,
	}
}

func (e *MessageEnricher) EnrichMessageWithChat(chat *models.Chat, msg *models.Message) (*models.Collection, error) {
	data := &models.Collection{}

	data.Chats[chat.ID] = chat
	data.Messages[msg.ID] = msg

	sender, err := e.userRepo.GetUserById(msg.SenderID)
	if err != nil {
		return nil, err
	}

	data.Users[sender.ID] = sender

	if msg.HasReply() {
		replyMessage, err := e.msgRepo.GetMessageById(msg.ReplyToId)
		if err != nil {
			return nil, err
		}
		data.Replies[replyMessage.ID] = replyMessage

		replySender, err := e.userRepo.GetUserById(replyMessage.SenderID)
		if err != nil {
			return nil, err
		}

		data.Users[replySender.ID] = replySender
	}

	return data, nil
}
