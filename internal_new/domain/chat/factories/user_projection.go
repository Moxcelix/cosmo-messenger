package factories

import (
	auth_models "main/internal_new/domain/auth/models"
	"main/internal_new/domain/chat/projections"
)

type UserProjectionFactory struct {
}

func NewUserProjectionFactory() *UserProjectionFactory {
	return &UserProjectionFactory{}
}

func (s *UserProjectionFactory) ProjectUser(user *auth_models.User) *projections.UserProjection {
	return &projections.UserProjection{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
}

func (s *UserProjectionFactory) ProjectUsers(users map[string]*auth_models.User) map[string]*projections.UserProjection {
	result := make(map[string]*projections.UserProjection, len(users))
	for id, user := range users {
		if user == nil {
			continue
		}
		result[id] = s.ProjectUser(user)
	}

	return result
}
