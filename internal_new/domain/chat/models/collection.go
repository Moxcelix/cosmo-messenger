package models

import (
	auth_models "main/internal_new/domain/auth/models"
)

type Collection struct {
	Chats    map[string]*Chat             `bson:"chats" json:"chats"`
	Messages map[string]*Message          `bson:"messages" json:"messages"`
	Replies  map[string]*Message          `bson:"replies" json:"replies"`
	Users    map[string]*auth_models.User `bson:"users" json:"users"`

	HasNext bool `bson:"has_next" json:"has_next"`
	HasPrev bool `bson:"has_prev" json:"has_prev"`
}
