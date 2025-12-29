package factory

import (
	"main/internal_new/domain/chat/errors"
	"main/internal_new/domain/chat/models"
	"time"
)

type ChatFactory struct {
}

func NewChatFactory() *ChatFactory {
	return &ChatFactory{}
}

func (f *ChatFactory) CreateDirectChat(user1ID, user2ID string) (*models.Chat, error) {
	if user1ID == user2ID {
		return nil, errors.ErrCannotCreateChatWithSelf
	}

	return &models.Chat{
		Type: models.ChatTypeDirect,
		Members: []*models.ChatMember{
			{UserID: user1ID, Role: models.RoleMember},
			{UserID: user2ID, Role: models.RoleMember},
		},
		CreatedAt: time.Now(),
	}, nil
}

func (f *ChatFactory) CreateGroupChat(
	creatorID, name string, initialMembers []string) (*models.Chat, error) {
	members := make([]*models.ChatMember, 0, len(initialMembers)+1)

	members = append(members, &models.ChatMember{
		UserID: creatorID,
		Role:   models.RoleAdmin,
	})

	for _, userID := range initialMembers {
		if userID == creatorID {
			continue
		}

		members = append(members, &models.ChatMember{
			UserID: userID,
			Role:   models.RoleMember,
		})
	}

	return &models.Chat{
		Type:      models.ChatTypeGroup,
		Name:      name,
		Members:   members,
		CreatedBy: creatorID,
	}, nil
}
