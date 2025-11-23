package usecases

import (
	auth_domain "main/internal/domain/auth"
	user_domain "main/internal/domain/user"
)

type ValidateUsecase struct {
	authservice auth_domain.AuthService
	userPolicy  *user_domain.UserPolicy
}

func NewValidateUsecase(
	authservice auth_domain.AuthService,
	userPolicy *user_domain.UserPolicy,
) *ValidateUsecase {
	return &ValidateUsecase{
		authservice: authservice,
		userPolicy:  userPolicy,
	}
}

func (uc *ValidateUsecase) Execute(accessToken string) (string, error) {
	userId, err := uc.authservice.ValidateAccessToken(accessToken)
	if err != nil {
		return "", err
	}

	if err := uc.userPolicy.ValidateExists(userId); err != nil {
		return "", err
	}

	return userId, nil
}
