package usecases

import (
	errors "main/internal_new/domain/auth/errors"
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

	login, err := uc.authservice.Login(user, password)
	if err != nil {
		return nil, err
	}

	return login, nil
}
