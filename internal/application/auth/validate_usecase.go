package auth_application

import (
	"main/internal/domain/auth"
	"main/internal/domain/user"
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

	user, err := uc.userReposiory.GetUserById(userId)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", user_domain.ErrUserNotFound
	}

	return user.Username, nil
}
