package usecases

import (
	auth_domain "main/internal/domain/auth"
	user_domain "main/internal/domain/user"
)

type LoginUsecase struct {
	authservice   auth_domain.AuthService
	userReposiory user_domain.UserRepository
}

func NewLoginUsecase(authservice auth_domain.AuthService, userReposiory user_domain.UserRepository) *LoginUsecase {
	return &LoginUsecase{
		authservice:   authservice,
		userReposiory: userReposiory,
	}
}

func (uc *LoginUsecase) Execute(username, password string) (string, string, error) {
	if user, _ := uc.userReposiory.GetUserByUsername(username); user == nil {
		return "", "", user_domain.ErrUserNotFound
	}

	return uc.authservice.Login(username, password)
}
