package chat_domain

import "time"

type ChatType string

const (
	ChatTypeDirect ChatType = "direct"
	ChatTypeGroup  ChatType = "group"
)

type ChatMemberRole string

const (
	RoleMember ChatMemberRole = "member"
	RoleAdmin  ChatMemberRole = "admin"
)

type Chat struct {
	ID          string       `json:"id" bson:"_id"`
	Type        ChatType     `json:"type" bson:"type"`
	Name        string       `json:"name" bson:"name"`
	Description string       `json:"description" bson:"description"`
	CreatedBy   string       `json:"created_by" bson:"created_by"`
	Members     []ChatMember `json:"members" bson:"members"`
	CreatedAt   time.Time    `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" bson:"updated_at"`
}

type ChatMember struct {
	UserID   string         `json:"user_id" bson:"user_id"`
	Role     ChatMemberRole `json:"role" bson:"role"`
	JoinedAt time.Time      `json:"joined_at" bson:"joined_at"`
}

type ChatList struct {
	Chats  []*Chat
	Total  int
	Offset int
	Limit  int
}
