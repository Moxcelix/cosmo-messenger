package models

import (
	"time"
)

type ChatMemberRole string

const (
	RoleMember ChatMemberRole = "member"
	RoleAdmin  ChatMemberRole = "admin"
)

type ChatMember struct {
	UserID   string         `json:"user_id" bson:"user_id"`
	Role     ChatMemberRole `json:"role" bson:"role"`
	JoinedAt time.Time      `json:"joined_at" bson:"joined_at"`
}
