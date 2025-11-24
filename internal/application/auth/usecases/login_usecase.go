package usecases

import (
	"main/internal/application/auth/dto"
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

func (uc *LoginUsecase) Execute(username, password string) (*dto.LoginData, error) {
	user, err := uc.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, user_domain.ErrUserNotFound
	}

	accessToken, refreshToken, err := uc.authservice.Login(user, password)
	if err != nil {
		return nil, err
	}

	return &dto.LoginData{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
