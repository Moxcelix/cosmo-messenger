package factories

import (
	auth_models "main/internal_new/domain/auth/models"
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

func (s *CollectionProjectionFactory) ProjectCollectionWithSingleChat(
	userId string,
	chat *models.Chat,
	messages map[string]*models.Message,
	replies map[string]*models.Message,
	users map[string]*auth_models.User,
	hasNext, hasPrev bool,
) (*projections.CollectionProjection, error) {

	chatsMap := make(map[string]*models.Chat)
	if chat != nil {
		chatsMap[chat.ID] = chat
	}

	return s.ProjectCollection(
		userId,
		chatsMap,
		messages,
		replies,
		users,
		hasNext,
		hasPrev,
	)
}

func (s *CollectionProjectionFactory) ProjectCollection(
	userId string,
	chats map[string]*models.Chat,
	messages map[string]*models.Message,
	replies map[string]*models.Message,
	users map[string]*auth_models.User,
	hasNext, hasPrev bool,
) (*projections.CollectionProjection, error) {
	messagesProjections := s.messageProjectionFactory.ProjectMessages(messages, DefaultMessage)
	repliesProjections := s.messageProjectionFactory.ProjectMessages(replies, ReplyOnly)
	chatsProjections := s.chatProjectionFactory.ProjectChats(chats, users, userId)
	usersProjections := s.userProjectionFactory.ProjectUsers(users)

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

		HasNext: hasNext,
		HasPrev: hasPrev,
	}, nil
}
