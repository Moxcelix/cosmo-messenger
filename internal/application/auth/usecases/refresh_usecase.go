package usecases

import (
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

func (uc *RefreshUsecase) Execute(refreshToken string) (string, error) {
	return uc.authservice.Refresh(refreshToken)
}
