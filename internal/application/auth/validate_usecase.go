package auth_application

import (
	auth_domain "main/internal/domain/auth"
	user_domain "main/internal/domain/user"
)

type ValidateUsecase struct {
	authservice   auth_domain.AuthService
	userReposiory user_domain.UserRepository
}

func NewValidateUsecase(
	authservice auth_domain.AuthService,
	userReposiory user_domain.UserRepository) *ValidateUsecase {
	return &ValidateUsecase{
		authservice:   authservice,
		userReposiory: userReposiory,
	}
}

func (uc *ValidateUsecase) Execute(accessToken string) (string, error) {
	userId, err := uc.authservice.ValidateAccessToken(accessToken)
	if err != nil {
		return "", err
	}

	exists, err := uc.userReposiory.UserExists(userId)
	if err != nil {
		return "", err
	}

	if !exists {
		return "", user_domain.ErrUserNotFound
	}

	return userId, nil
}
