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

	messageList, err := uc.msgRepo.GetMessagesByChatId(
		chatId, cursorMessageId, count, direction, message_domain.FullPopulate)
	if err != nil {
		return nil, err
	}

	return uc.mapToChatMessages(messageList), nil
}

func (uc *GetChatMessagesUsecase) mapToChatMessages(domainList *message_domain.MessageList) *ChatMessages {
	if domainList == nil {
		return &ChatMessages{
			Messages: []*ChatMessage{},
			Meta:     ScrollingMeta{},
		}
	}

	messages := make([]*ChatMessage, 0, len(domainList.Messages))
	for _, domainMsg := range domainList.Messages {
		messages = append(messages, uc.mapToChatMessage(domainMsg))
	}

	hasNext := domainList.Offset+len(messages) < domainList.Total
	hasPrev := domainList.Offset > 0

	return &ChatMessages{
		Messages: messages,
		Meta: ScrollingMeta{
			HasPrev: hasPrev,
			HasNext: hasNext,
			Offset:  domainList.Offset,
			Total:   domainList.Total,
		},
	}
}

func (uc *GetChatMessagesUsecase) mapToChatMessage(domainMsg *message_domain.Message) *ChatMessage {
	if domainMsg == nil {
		return nil
	}

	chatMessage := &ChatMessage{
		ID:        domainMsg.ID,
		Content:   domainMsg.Content,
		Timestamp: domainMsg.CreatedAt,
		Edited:    !domainMsg.UpdatedAt.Equal(domainMsg.CreatedAt),
	}

	// Map sender
	if domainMsg.Sender != nil {
		chatMessage.Sender = &Sender{
			ID:   domainMsg.SenderID,
			Name: domainMsg.Sender.Name,
		}
	} else {
		// Fallback - create minimal sender from ID
		chatMessage.Sender = &Sender{
			ID:   domainMsg.SenderID,
			Name: "Unknown", // или можно получить из userRepo если нужно
		}
	}

	// Map reply
	if domainMsg.Replied != nil {
		reply := &RepliedMessage{
			ID:      domainMsg.ReplyToID,
			Content: domainMsg.Replied.Content,
		}

		if domainMsg.Replied.Sender != nil {
			reply.Sender = &Sender{
				ID:   domainMsg.Replied.SenderID,
				Name: domainMsg.Replied.Sender.Name,
			}
		} else {
			// Fallback for reply sender
			reply.Sender = &Sender{
				ID:   domainMsg.ReplyToID, // используем ID из основного сообщения
				Name: "Unknown",
			}
		}

		chatMessage.ReplyTo = reply
	}

	return chatMessage
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
