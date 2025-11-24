package usecases

import (
	"main/internal/application/auth/dto"
	auth "main/internal/domain/auth"
)

type RefreshUsecase struct {
	authservice auth.AuthService
}

func NewRefreshUsecase(authservice auth.AuthService) *RefreshUsecase {
	return &RefreshUsecase{
		authservice: authservice,
	}
}

func (uc *RefreshUsecase) Execute(refreshToken string) (*dto.RefreshData, error) {
	accessToken, err := uc.authservice.Refresh(refreshToken)
	if err != nil {
		return nil, err
	}

	return &dto.RefreshData{
		AccessToken: accessToken,
	}, nil
}
