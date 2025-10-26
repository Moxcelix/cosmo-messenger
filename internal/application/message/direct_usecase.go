package message_application

import (
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
)

type DirectMessageUsecase struct {
	chatProvider  *chat_domain.DirectChatProvider
	messagePolicy *message_domain.MessagePolicy
	messageRepo   message_domain.MessageRepository
}

func NewDirectMessageUsecase(
	chatProvider *chat_domain.DirectChatProvider,
	messagePolicy *message_domain.MessagePolicy,
	messageRepo message_domain.MessageRepository,
) *DirectMessageUsecase {
	return &DirectMessageUsecase{
		chatProvider:  chatProvider,
		messageRepo:   messageRepo,
		messagePolicy: messagePolicy,
	}
}

func (uc *DirectMessageUsecase) Execute(senderId, receiverId, content string) error {
	chat, err := uc.chatProvider.Provide(senderId, receiverId)
	if err != nil {
		return err
	}

	if err := uc.messagePolicy.ValidateMessageContent(content); err != nil {
		return err
	}

	message := &message_domain.Message{
		ChatID:   chat.ID,
		SenderID: senderId,
		Content:  content,
	}

	return uc.messageRepo.CreateMessage(message)
}
