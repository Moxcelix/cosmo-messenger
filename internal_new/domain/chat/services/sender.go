package services

import (
	auth_models "main/internal_new/domain/auth/models"
	"main/internal_new/domain/chat/projections"
)

type SenderService struct {
}

func NewSenderService() *SenderService {
	return &SenderService{}
}

func (s *SenderService) ProjectUser(user *auth_models.User) *projections.Sender {
	return &projections.Sender{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
}

func (s *SenderService) ProjectUsers(users map[string]*auth_models.User) map[string]*projections.Sender {
	result := make(map[string]*projections.Sender, len(users))
	for id, user := range users {
		if user == nil {
			continue
		}
		result[id] = s.ProjectUser(user)
	}

	return result
}
