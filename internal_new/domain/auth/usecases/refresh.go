package usecases

import (
	"main/internal_new/domain/auth/models"
	"main/internal_new/domain/auth/services"
)

type RefreshUsecase struct {
	authservice services.AuthService
}

func NewRefreshUsecase(authservice services.AuthService) *RefreshUsecase {
	return &RefreshUsecase{
		authservice: authservice,
	}
}

func (uc *RefreshUsecase) Execute(refreshToken string) (*models.Refresh, error) {
	accessToken, err := uc.authservice.Refresh(refreshToken)
	if err != nil {
		return nil, err
	}
	// TODO: transport to service
	return &models.Refresh{
		AccessToken: accessToken,
	}, nil
}
