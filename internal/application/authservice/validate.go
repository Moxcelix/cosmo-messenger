package authservice_application

import (
	"main/internal/domain/authservice"
)

type ValidateUsecase struct {
	authservice authservice.AuthService
}

func NewValidateUsecase(authservice authservice.AuthService) *ValidateUsecase {
	return &ValidateUsecase{
		authservice: authservice,
	}
}

func (uc *ValidateUsecase) Execute(accessToken string) (string, error) {
	return uc.authservice.ValidateAccessToken(accessToken)
}
