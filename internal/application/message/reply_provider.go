package message_application

import (
	user_application "main/internal/application/user"
	message_domain "main/internal/domain/message"
)

type ReplyProvider struct {
	msgRepo        message_domain.MessageRepository
	senderProvider *user_application.SenderProvider
}

func NewReplyProvider(
	msgRepo message_domain.MessageRepository,
	senderProvider *user_application.SenderProvider,
) *ReplyProvider {
	return &ReplyProvider{
		msgRepo:        msgRepo,
		senderProvider: senderProvider,
	}
}

func (p *ReplyProvider) Provide(msgId string) (*Reply, error) {
	msg, err := p.msgRepo.GetMessageById(msgId)

	if err != nil {
		return nil, err
	}

	if msg == nil {
		return nil, nil
	}

	sender, err := p.senderProvider.Provide(msg.SenderID)
	if err != nil {
		return nil, err
	}

	return &Reply{
		ID:      msg.ID,
		Content: msg.Content,
		Sender:  sender,
	}, nil
}
