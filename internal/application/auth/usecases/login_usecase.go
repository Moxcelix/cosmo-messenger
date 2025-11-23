package usecases

import (
	auth_domain "main/internal/domain/auth"
	user_domain "main/internal/domain/user"
)

type LoginUsecase struct {
	authservice auth_domain.AuthService
	userRepo    user_domain.UserRepository
}

func NewLoginUsecase(
	authservice auth_domain.AuthService,
	userRepo user_domain.UserRepository,
) *LoginUsecase {
	return &LoginUsecase{
		authservice: authservice,
		userRepo:    userRepo,
	}
}

func (uc *LoginUsecase) Execute(username, password string) (string, string, error) {
	user, err := uc.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", "", err
	}
	if user == nil {
		return "", "", user_domain.ErrUserNotFound
	}

	return uc.authservice.Login(username, password)
}
