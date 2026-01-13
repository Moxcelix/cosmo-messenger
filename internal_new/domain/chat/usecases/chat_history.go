package usecases

import (
	"main/internal_new/domain/chat/factories"
	"main/internal_new/domain/chat/projections"
	"main/internal_new/domain/chat/queries"
	"main/internal_new/domain/chat/services"
)

type ChatHistoryUsecase struct {
	userChatService             *services.UserChatService
	messageCollectionQuery      queries.MessageCollectionQuery
	collectionProjectionFactory *factories.CollectionProjectionFactory
}

func NewChatHistoryUsecase(
	userChatService *services.UserChatService,
	messageCollectionQuery queries.MessageCollectionQuery,
	collectionProjectionFactory *factories.CollectionProjectionFactory,
) *ChatHistoryUsecase {
	return &ChatHistoryUsecase{
		userChatService:             userChatService,
		messageCollectionQuery:      messageCollectionQuery,
		collectionProjectionFactory: collectionProjectionFactory,
	}
}

func (uc *ChatHistoryUsecase) Execute(
	userId, chatId, cursorMessageId string, count int, direction string) (*projections.CollectionProjection, error) {

	chat, err := uc.userChatService.GetChatForUser(chatId, userId)
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
