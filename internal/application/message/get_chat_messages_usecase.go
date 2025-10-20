package message_application

import (
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	user_domain "main/internal/domain/user"
)

const (
	defaultCount = 10
	maxPageSize  = 100
)

type GetChatMessagesUsecase struct {
	chatRepo chat_domain.ChatRepository
	msgRepo  message_domain.MessageRepository
	userRepo user_domain.UserRepository
}

func NewGetChatMessagesUsecase(
	chatRepo chat_domain.ChatRepository,
	msgRepo message_domain.MessageRepository,
	userRepo user_domain.UserRepository) *GetChatMessagesUsecase {
	return &GetChatMessagesUsecase{
		chatRepo: chatRepo,
		msgRepo:  msgRepo,
		userRepo: userRepo,
	}
}

func (uc *GetChatMessagesUsecase) Execute(
	userId, chatId, cursorMessageId string, count int, direction string) (*ChatMessages, error) {
	if err := uc.validateChat(chatId); err != nil {
		return nil, err
	}

	if err := uc.validateUserAccess(userId, chatId); err != nil {
		return nil, err
	}

	if count < 1 {
		count = defaultCount
	}
	if count > maxPageSize {
		count = maxPageSize
	}

	messageList, err := uc.msgRepo.GetMessagesByChatIdScroll(chatId, cursorMessageId, count, direction)
	if err != nil {
		return nil, err
	}

	return uc.toResponseModel(messageList)
}

func (uc *GetChatMessagesUsecase) toResponseModel(messageList *message_domain.MessageList) (*ChatMessages, error) {
	messages := make([]*ChatMessage, 0, len(messageList.Messages))
	for _, msg := range messageList.Messages {
		sender, err := uc.getSenderModel(msg.SenderID)
		if err != nil {
			return nil, err
		}
		repliedMessage, err := uc.getRepliedMessage(msg.ReplyTo)
		if err != nil {
			return nil, err
		}
		timestamp := msg.CreatedAt
		edited := !msg.UpdatedAt.Equal(msg.CreatedAt)

		message := &ChatMessage{
			ID:        msg.ID,
			Content:   msg.Content,
			ReplyTo:   repliedMessage,
			Sender:    sender,
			Timestamp: timestamp,
			Edited:    edited,
		}

		messages = append(messages, message)
	}
	chatMessages := &ChatMessages{
		Messages: messages,
		Meta: ScrollingMeta{
			HasPrev: messageList.Offset > 0,
			HasNext: messageList.Offset < messageList.Total-messageList.Limit,
			Offset:  messageList.Offset,
			Total:   messageList.Total,
		},
	}

	return chatMessages, nil
}

func (uc *GetChatMessagesUsecase) getRepliedMessage(msgId string) (*RepliedMessage, error) {
	message, err := uc.msgRepo.GetMessageById(msgId)
	if err == nil {
		return nil, err
	}

	if message == nil {
		return nil, nil
	}

	sender, err := uc.getSenderModel(message.SenderID)
	if err == nil {
		return nil, err
	}

	return &RepliedMessage{
		ID:      message.ID,
		Content: message.Content,
		Sender:  sender,
	}, nil
}

func (uc *GetChatMessagesUsecase) getSenderModel(senderId string) (*Sender, error) {
	sender, err := uc.userRepo.GetUserById(senderId)
	if err != nil {
		return nil, err
	}

	if sender == nil {
		return nil, nil
	}

	return &Sender{
		ID:   sender.ID,
		Name: sender.Name,
	}, nil
}

func (uc *GetChatMessagesUsecase) validateChat(chatId string) error {
	exists, err := uc.chatRepo.ChatExists(chatId)
	if err != nil {
		return err
	}
	if !exists {
		return chat_domain.ErrChatNotFound
	}
	return nil
}

func (uc *GetChatMessagesUsecase) validateUserAccess(userId, chatId string) error {
	hasAccess, err := uc.chatRepo.UserInChat(userId, chatId)
	if err != nil {
		return err
	}
	if !hasAccess {
		return chat_domain.ErrChatAccessDenied
	}
	return nil
}
