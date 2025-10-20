package chat_application

import (
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	user_domain "main/internal/domain/user"
)

const (
	defaultPage  = 1
	defaultCount = 10
	maxPageSize  = 100
)

type GetUserChatsUsecase struct {
	chatRepo chat_domain.ChatRepository
	msgRepo  message_domain.MessageRepository
	userRepo user_domain.UserRepository
}

func NewGetUserChatsUsecase(
	chatRepo chat_domain.ChatRepository,
	msgRepo message_domain.MessageRepository,
	userRepo user_domain.UserRepository) *GetUserChatsUsecase {
	return &GetUserChatsUsecase{
		chatRepo: chatRepo,
		msgRepo:  msgRepo,
		userRepo: userRepo,
	}
}

func (uc *GetUserChatsUsecase) Execute(userID string, page, count int) (*UserChats, error) {
	if page < 1 {
		page = defaultPage
	}
	if count < 1 {
		count = defaultCount
	}
	if count > maxPageSize {
		count = maxPageSize
	}

	offset := (page - 1) * count

	chatList, err := uc.chatRepo.GetUserChats(userID, offset, count)
	if err != nil {
		return nil, err
	}

	totalPages := 0
	if chatList.Total > 0 {
		totalPages = (chatList.Total + count - 1) / count
	}

	meta := PaginationMeta{
		HasPrev:    page > 1,
		HasNext:    page < totalPages,
		Page:       page,
		TotalPages: totalPages,
		Total:      chatList.Total,
		Count:      len(chatList.Chats),
	}

	chats := make([]*ChatListItem, 0, len(chatList.Chats))
	for _, chat := range chatList.Chats {
		chatResponse := &ChatListItem{
			ID:   chat.ID,
			Name: chat.Name,
			Type: chat.Type,
		}

		lastMessage, err := uc.msgRepo.GetLastChatMessage(chat.ID)
		if err != nil {
			return nil, err
		}

		if lastMessage != nil {
			timestamp := lastMessage.CreatedAt
			if !lastMessage.UpdatedAt.Equal(lastMessage.CreatedAt) {
				timestamp = lastMessage.UpdatedAt
			}
			chatResponse.LastMessage = &LastMessage{
				ID:        lastMessage.ID,
				Content:   lastMessage.Content,
				Timestamp: timestamp,
			}

			sender, err := uc.userRepo.GetUserById(lastMessage.SenderID)
			if err != nil {
				return nil, err
			}

			if sender != nil {
				chatResponse.LastMessage.Sender = &Sender{
					ID:   sender.ID,
					Name: sender.Name,
				}
			}
		}

		chats = append(chats, chatResponse)
	}

	return &UserChats{
		Chats: chats,
		Meta:  meta,
	}, nil
}
