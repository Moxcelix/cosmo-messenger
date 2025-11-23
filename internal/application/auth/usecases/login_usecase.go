package usecases

import (
	auth_domain "main/internal/domain/auth"
	user_domain "main/internal/domain/user"
)

type LoginUsecase struct {
	authservice auth_domain.AuthService
	userPolicy  *user_domain.UserPolicy
}

func NewLoginUsecase(
	authservice auth_domain.AuthService,
	userPolicy *user_domain.UserPolicy,
) *LoginUsecase {
	return &LoginUsecase{
		authservice: authservice,
		userPolicy:  userPolicy,
	}
}

func (uc *LoginUsecase) Execute(username, password string) (string, string, error) {
	if err := uc.userPolicy.ValidateExistsByUsername(username); err != nil {
		return "", "", err
	}

	return uc.authservice.Login(username, password)
}
