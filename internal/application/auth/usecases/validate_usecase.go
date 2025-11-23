package usecases

import (
	auth_domain "main/internal/domain/auth"
	user_domain "main/internal/domain/user"
)

type ValidateUsecase struct {
	authservice auth_domain.AuthService
	userRepo    user_domain.UserRepository
}

func NewValidateUsecase(
	authservice auth_domain.AuthService,
	userRepo user_domain.UserRepository,
) *ValidateUsecase {
	return &ValidateUsecase{
		authservice: authservice,
		userRepo:    userRepo,
	}
}

func (uc *ValidateUsecase) Execute(accessToken string) (string, error) {
	userId, err := uc.authservice.ValidateAccessToken(accessToken)
	if err != nil {
		return "", err
	}

	user, err := uc.userRepo.GetUserById(userId)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", user_domain.ErrUserNotFound
	}

	return userId, nil
}
