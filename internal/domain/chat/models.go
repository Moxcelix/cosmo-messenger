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
	ID          string        `json:"id" bson:"_id"`
	Type        ChatType      `json:"type" bson:"type"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	CreatedBy   string        `json:"created_by" bson:"created_by"`
	Members     []*ChatMember `json:"members" bson:"members"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at"`
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

func (c *Chat) GetMembersId() []string {
	memberIDs := make([]string, 0, len(c.Members))
	for _, member := range c.Members {
		if member != nil && member.UserID != "" {
			memberIDs = append(memberIDs, member.UserID)
		}
	}
	return memberIDs
}

func (c *Chat) GetAdminIds() []string {
	adminIDs := make([]string, 0)
	for _, member := range c.Members {
		if member != nil && member.UserID != "" && member.Role == RoleAdmin {
			adminIDs = append(adminIDs, member.UserID)
		}
	}
	return adminIDs
}

func (c *Chat) GetMemberIdsExcluding(excludeUserID string) []string {
	memberIDs := make([]string, 0, len(c.Members))
	for _, member := range c.Members {
		if member != nil && member.UserID != "" && member.UserID != excludeUserID {
			memberIDs = append(memberIDs, member.UserID)
		}
	}
	return memberIDs
}

func (c *Chat) GetMemberIdsByRole(role ChatMemberRole) []string {
	memberIDs := make([]string, 0)
	for _, member := range c.Members {
		if member != nil && member.UserID != "" && member.Role == role {
			memberIDs = append(memberIDs, member.UserID)
		}
	}
	return memberIDs
}

func (c *Chat) HasMember(userID string) bool {
	for _, member := range c.Members {
		if member != nil && member.UserID == userID {
			return true
		}
	}
	return false
}

func (c *Chat) GetMemberRole(userID string) (ChatMemberRole, bool) {
	for _, member := range c.Members {
		if member != nil && member.UserID == userID {
			return member.Role, true
		}
	}
	return "", false
}
