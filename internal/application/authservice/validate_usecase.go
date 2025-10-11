package authservice_application

import (
	"main/internal/domain/authservice"
	"main/internal/domain/userservice"
)

type ValidateUsecase struct {
	authservice   authservice.AuthService
	userReposiory userservice.UserRepository
}

func NewValidateUsecase(
	authservice authservice.AuthService,
	userReposiory userservice.UserRepository) *ValidateUsecase {
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
		return "", userservice.ErrUserNotFound
	}

	return user.Username, nil
}
