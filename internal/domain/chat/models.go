package chat_domain

import "time"

type ChatMember struct {
	UserID string `json:"user_id" bson:"user_id"`
	Role   string `json:"role" bson:"role"`
}

type Chat struct {
	ID        string       `json:"id" bson:"_id"`
	Name      string       `json:"name" bson:"name"`
	OwnerID   string       `json:"owner_id" bson:"owner_id"`
	Members   []ChatMember `json:"members" bson:"members"`
	CreatedAt time.Time    `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time    `bson:"updated_at" json:"updated_at"`
}
