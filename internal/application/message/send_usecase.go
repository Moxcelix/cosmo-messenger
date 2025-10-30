package message_application

import (
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	"time"
)

type SendMessageUsecase struct {
	msgRepo            message_domain.MessageRepository
	chatRepo           chat_domain.ChatRepository
	messagePolicy      *message_domain.MessagePolicy
	messageAssembler   *ChatMessageAssembler
	messageBroadcaster MessageBroadcaster
}

func NewSendMessageUsecase(
	msgRepo message_domain.MessageRepository,
	chatRepo chat_domain.ChatRepository,
	messagePolicy *message_domain.MessagePolicy,
	messageAssembler *ChatMessageAssembler,
	messageBroadcaster MessageBroadcaster,
) *SendMessageUsecase {
	return &SendMessageUsecase{
		msgRepo:            msgRepo,
		chatRepo:           chatRepo,
		messagePolicy:      messagePolicy,
		messageAssembler:   messageAssembler,
		messageBroadcaster: messageBroadcaster,
	}
}

func (uc *SendMessageUsecase) Execute(senderId, chatId, content string) error {
	chat, err := uc.chatRepo.GetByID(chatId)
	if err != nil {
		return err
	}

	if chat == nil {
		return chat_domain.ErrChatNotFound
	}

	if err := uc.messagePolicy.ValidateMessageContent(content); err != nil {
		return err
	}

	message := &message_domain.Message{
		ChatID:   chatId,
		SenderID: senderId,
		Content:  content,
	}

	if err := uc.msgRepo.CreateMessage(message); err != nil {
		return err
	}

	if err := uc.chatRepo.MarkUpdated(chatId, time.Now()); err != nil {
		return err
	}

	msgDto, err := uc.messageAssembler.Assemble(message)
	if err != nil {
		return err
	}

	chatMembersId := chat.GetMembersId()
	if err := uc.messageBroadcaster.BroadcastToUsers(chatMembersId, msgDto); err != nil {
		return err
	}

	return nil
}
