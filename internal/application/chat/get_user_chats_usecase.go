package chat_application

import chat_domain "main/internal/domain/chat"

const (
	defaultPage  = 1
	defaultCount = 10
	maxPageSize  = 100
)

type GetUserChatsUsecase struct {
	chatRepo chat_domain.ChatRepository
}

func NewGetUserChatsUsecase(chatRepo chat_domain.ChatRepository) *GetUserChatsUsecase {
	return &GetUserChatsUsecase{
		chatRepo: chatRepo,
	}
}

func (uc *GetUserChatsUsecase) Execute(
	userId string, page, count int) (*chat_domain.ChatList, error) {
	if page < 1 {
		page = defaultPage
	}
	if count < 1 {
		count = defaultCount
	}
	if count > maxPageSize {
		count = maxPageSize
	}

	return uc.chatRepo.GetUserChats(userId, (page-1)*count, count)
}
