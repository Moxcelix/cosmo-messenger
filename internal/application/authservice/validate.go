package authservice_application

import (
	authservice "main/internal/domain/authservice"
)

type ValidateUsecase struct {
	authservice authservice.AuthService
}

func (uc *ValidateUsecase) Execute(accessToken string) (string, error) {
	return uc.authservice.ValidateAccessToken(accessToken)
}
