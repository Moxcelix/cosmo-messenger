package usecases

import (
	"main/internal/application/auth/dto"
	auth_domain "main/internal/domain/auth"
	user_domain "main/internal/domain/user"
)

type ValidateUsecase struct {
	authservice auth_domain.AuthService
	userRepo    user_domain.UserRepository
}

func NewValidateUsecase(
	authservice auth_domain.AuthService,
	userRepo user_domain.UserRepository,
) *ValidateUsecase {
	return &ValidateUsecase{
		authservice: authservice,
		userRepo:    userRepo,
	}
}

func (uc *ValidateUsecase) Execute(accessToken string) (*dto.ValidateData, error) {
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

	return &dto.ValidateData{
		UserID: userId,
	}, nil
}
