package mappers

import (
	"main/internal/application/message/dto"
	user_application "main/internal/application/user"
	message_domain "main/internal/domain/message"
)

type ChatMessageAssembler struct {
	replyProvider  *ReplyProvider
	senderProvider *user_application.SenderProvider
}

func NewChatMessageAssembler(
	replyProvider *ReplyProvider,
	senderProvider *user_application.SenderProvider,
) *ChatMessageAssembler {
	return &ChatMessageAssembler{
		replyProvider:  replyProvider,
		senderProvider: senderProvider,
	}
}

func (a *ChatMessageAssembler) Assemble(msg *message_domain.Message) (*dto.ChatMessage, error) {
	sender, err := a.senderProvider.Provide(msg.SenderID)
	if err != nil {
		return nil, err
	}
	repliedMessage, err := a.replyProvider.Provide(msg.ReplyTo)
	if err != nil {
		return nil, err
	}
	timestamp := msg.CreatedAt
	edited := !msg.UpdatedAt.Equal(msg.CreatedAt)

	return &dto.ChatMessage{
		ID:        msg.ID,
		ChatID:    msg.ChatID,
		Content:   msg.Content,
		ReplyTo:   repliedMessage,
		Sender:    sender,
		Timestamp: timestamp,
		Edited:    edited,
	}, nil
}
