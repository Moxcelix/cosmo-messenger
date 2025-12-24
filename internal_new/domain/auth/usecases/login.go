package usecases

import (
	errors "main/internal_new/domain/auth/erorrs"
	"main/internal_new/domain/auth/models"
	"main/internal_new/domain/auth/repositories"
	"main/internal_new/domain/auth/services"
)

type LoginUsecase struct {
	authservice services.AuthService
	userRepo    repositories.UserRepository
}

func NewLoginUsecase(
	authservice services.AuthService,
	userRepo repositories.UserRepository,
) *LoginUsecase {
	return &LoginUsecase{
		authservice: authservice,
		userRepo:    userRepo,
	}
}

func (uc *LoginUsecase) Execute(username, password string) (*models.Login, error) {
	user, err := uc.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.ErrUserNotFound
	}

	accessToken, refreshToken, err := uc.authservice.Login(user, password)
	if err != nil {
		return nil, err
	}

	// TODO: transport to service
	return &models.Login{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
