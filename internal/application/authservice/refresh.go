package authservice_application

import (
	authservice "main/internal/domain/authservice"
)

type RefreshUsecase struct {
	authservice authservice.AuthService
}

func (uc *RefreshUsecase) Execute(refreshToken string) (string, error) {
	return uc.authservice.Refresh(refreshToken)
}
