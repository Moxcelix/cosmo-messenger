package queries

import (
	auth_models "main/internal_new/domain/auth/models"
	"main/internal_new/domain/chat/models"
)

type MessageCollection struct {
	Messages map[string]*models.Message   `bson:"messages" json:"messages"`
	Replies  map[string]*models.Message   `bson:"replies" json:"replies"`
	Users    map[string]*auth_models.User `bson:"users" json:"users"`

	HasNext bool `bson:"has_next" json:"has_next"`
	HasPrev bool `bson:"has_prev" json:"has_prev"`
}

type MessageCollectionQuery interface {
	Query(
		chatId string,
		cursorMessageId string,
		count int,
		direction string,
	) (
		collection *MessageCollection,
		err error,
	)
}
