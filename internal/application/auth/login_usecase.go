package authservice_application

import (
	"main/internal/domain/authservice"
	"main/internal/domain/userservice"
)

type LoginUsecase struct {
	authservice   authservice.AuthService
	userReposiory userservice.UserRepository
}

func NewLoginUsecase(authservice authservice.AuthService, userReposiory userservice.UserRepository) *LoginUsecase {
	return &LoginUsecase{
		authservice:   authservice,
		userReposiory: userReposiory,
	}
}

func (uc *LoginUsecase) Execute(username, password string) (string, string, error) {
	if user, _ := uc.userReposiory.GetUserByUsername(username); user == nil {
		return "", "", userservice.ErrUserNotFound
	}

	return uc.authservice.Login(username, password)
}
