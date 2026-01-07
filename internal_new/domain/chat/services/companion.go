package services

import (
	auth_errors "main/internal_new/domain/auth/errors"
	auth_models "main/internal_new/domain/auth/models"
	auth_repositories "main/internal_new/domain/auth/repositories"
)

type CompanionService struct {
	userRepo auth_repositories.UserRepository
}

func NewCompanionService(userRepo auth_repositories.UserRepository) *CompanionService {
	return &CompanionService{
		userRepo: userRepo,
	}
}

// TODO: complex access verification
func (s *CompanionService) GetCompanion(requesterId string, companionUsername string) (*auth_models.User, error) {
	companion, err := s.userRepo.GetUserByUsername(companionUsername)
	if err != nil {
		return nil, err
	}

	if companion == nil {
		return nil, auth_errors.ErrUserNotFound
	}

	return companion, nil
}
