package usecases

import (
	"main/internal_new/domain/chat/factories"
	"main/internal_new/domain/chat/projections"
	"main/internal_new/domain/chat/queries"
	"main/internal_new/domain/chat/services"
)

type DirectHistoryUsecase struct {
	messageCollectionQuery      queries.MessageCollectionQuery
	companionService            *services.CompanionService
	directService               *services.DirectChatService
	collectionProjectionFactory *factories.CollectionProjectionFactory
}

func NewDirectHistoryUsecase(
	messageCollectionQuery queries.MessageCollectionQuery,
	companionService *services.CompanionService,
	directService *services.DirectChatService,
	collectionProjectionFactory *factories.CollectionProjectionFactory,
) *DirectHistoryUsecase {
	return &DirectHistoryUsecase{
		messageCollectionQuery:      messageCollectionQuery,
		companionService:            companionService,
		directService:               directService,
		collectionProjectionFactory: collectionProjectionFactory,
	}
}

func (uc *DirectHistoryUsecase) Execute(
	userId, targetUsername, cursorMessageId string, count int, direction string) (*projections.CollectionProjection, error) {

	companion, err := uc.companionService.GetCompanion(userId, targetUsername)
	if err != nil {
		return nil, err
	}

	chat, err := uc.directService.GetDirectChat(userId, companion.ID)
	if err != nil {
		return nil, err
	}

	collection, err := uc.messageCollectionQuery.Query(chat.ID, cursorMessageId, count, direction)
	if err != nil {
		return nil, err
	}

	history, err := uc.collectionProjectionFactory.ProjectCollection(userId, collection)
	if err != nil {
		return nil, err
	}

	return history, nil
}
