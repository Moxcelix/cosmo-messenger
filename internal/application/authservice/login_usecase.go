package authservice_application

import (
	"main/internal/domain/authservice"
)

type LoginUsecase struct {
	authservice authservice.AuthService
}

func NewLoginUsecase(authservice authservice.AuthService) *LoginUsecase {
	return &LoginUsecase{
		authservice: authservice,
	}
}

func (uc *LoginUsecase) Execute(username, password string) (string, string, error) {
	return uc.authservice.Login(username, password)
}
