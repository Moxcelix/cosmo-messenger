package usecases

import (
	"main/internal/application/chat/dto"
	"main/internal/application/chat/services"
	chat_domain "main/internal/domain/chat"
)

const (
	defaultPage  = 1
	defaultCount = 10
	maxPageSize  = 100
)

type GetUserChatsUsecase struct {
	chatRepo            chat_domain.ChatRepository
	collectionAssembler *services.ChatCollectionAssembler
}

func NewGetUserChatsUsecase(
	chatRepo chat_domain.ChatRepository,
	collectionAssembler *services.ChatCollectionAssembler,
) *GetUserChatsUsecase {
	return &GetUserChatsUsecase{
		chatRepo:            chatRepo,
		collectionAssembler: collectionAssembler,
	}
}

func (uc *GetUserChatsUsecase) Execute(userID string, page, count int) (*dto.ChatCollection, error) {
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

	return uc.collectionAssembler.Assemble(chatList, userID)
}
