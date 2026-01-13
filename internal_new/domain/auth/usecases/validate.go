package usecases

import (
	"main/internal_new/domain/auth/models"
	"main/internal_new/domain/auth/services"
)

type ValidateUsecase struct {
	authservice services.AuthService
}

func NewValidateUsecase(
	authservice services.AuthService,
) *ValidateUsecase {
	return &ValidateUsecase{
		authservice: authservice,
	}
}

func (uc *ValidateUsecase) Execute(accessToken string) (*models.Validate, error) {
	vaildate, err := uc.authservice.ValidateAccessToken(accessToken)
	if err != nil {
		return nil, err
	}

	return vaildate, nil
}
