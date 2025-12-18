package usecases

import (
	user_domain "main/internal/domain/user"
	"main/internal_new/domain/auth/models"
	"main/internal_new/domain/auth/repositories"
	"main/internal_new/domain/auth/services"
)

type ValidateUsecase struct {
	authservice services.AuthService
	userRepo    repositories.UserRepository
}

func NewValidateUsecase(
	authservice services.AuthService,
	userRepo repositories.UserRepository,
) *ValidateUsecase {
	return &ValidateUsecase{
		authservice: authservice,
		userRepo:    userRepo,
	}
}

func (uc *ValidateUsecase) Execute(accessToken string) (*models.Validate, error) {
	userId, err := uc.authservice.ValidateAccessToken(accessToken)
	if err != nil {
		return nil, err
	}

	user, err := uc.userRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, user_domain.ErrUserNotFound
	}

	return &models.Validate{
		UserID: userId,
	}, nil
}
