package usecases

import (
	"main/internal_new/domain/chat/factories"
	"main/internal_new/domain/chat/projections"
	"main/internal_new/domain/chat/queries"
)

const (
	defaultCount = 10
	maxPageSize  = 100
)

type ChatCollectionUsecase struct {
	chatCollectionQuery         queries.ChatCollectionQuery
	collectionProjectionFactory *factories.CollectionProjectionFactory
}

func NewChatCollectionUsecase(
	chatCollectionQuery queries.ChatCollectionQuery,
	collectionProjectionFactory *factories.CollectionProjectionFactory,
) *ChatCollectionUsecase {
	return &ChatCollectionUsecase{
		chatCollectionQuery:         chatCollectionQuery,
		collectionProjectionFactory: collectionProjectionFactory,
	}
}

func (uc *ChatCollectionUsecase) Execute(
	userID string, cursorChatID string, count int, direction string) (*projections.CollectionProjection, error) {

	if count < 1 {
		count = defaultCount
	}
	if count > maxPageSize {
		count = maxPageSize
	}

	collection, err := uc.chatCollectionQuery.Query(userID, cursorChatID, count, direction)
	if err != nil {
		return nil, err
	}

	chatsCollection, err := uc.collectionProjectionFactory.ProjectCollection(userID, collection)
	if err != nil {
		return nil, err
	}

	return chatsCollection, nil
}
