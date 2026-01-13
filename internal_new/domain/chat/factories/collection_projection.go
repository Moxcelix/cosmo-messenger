package factories

import (
	"main/internal_new/domain/chat/models"
	"main/internal_new/domain/chat/projections"
)

type CollectionProjectionFactory struct {
	messageProjectionFactory *MessageProjectionFactory
	chatProjectionFactory    *ChatProjectionFactory
	userProjectionFactory    *UserProjectionFactory
}

func NewCollectionProjectionFactory(
	messageProjectionFactory *MessageProjectionFactory,
	chatProjectionFactory *ChatProjectionFactory,
	userProjectionFactory *UserProjectionFactory,
) *CollectionProjectionFactory {
	return &CollectionProjectionFactory{
		messageProjectionFactory: messageProjectionFactory,
		chatProjectionFactory:    chatProjectionFactory,
		userProjectionFactory:    userProjectionFactory,
	}
}

func (s *CollectionProjectionFactory) ProjectCollection(
	userId string, collection *models.Collection) (*projections.CollectionProjection, error) {

	messagesProjections := s.messageProjectionFactory.ProjectMessages(collection.Messages, DefaultMessage)
	repliesProjections := s.messageProjectionFactory.ProjectMessages(collection.Replies, ReplyOnly)
	chatsProjections := s.chatProjectionFactory.ProjectChats(collection.Chats, collection.Users, userId)
	usersProjections := s.userProjectionFactory.ProjectUsers(collection.Users)

	for id, reply := range repliesProjections {
		if _, exists := messagesProjections[id]; exists {
			continue
		}
		messagesProjections[id] = reply
	}

	return &projections.CollectionProjection{
		Chats:    chatsProjections,
		Messages: messagesProjections,
		Users:    usersProjections,

		HasNext: collection.HasNext,
		HasPrev: collection.HasPrev,
	}, nil
}
