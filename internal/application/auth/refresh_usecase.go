package authservice_application

import (
	"main/internal/domain/authservice"
)

type RefreshUsecase struct {
	authservice authservice.AuthService
}

func NewRefreshUsecase(authservice authservice.AuthService) *RefreshUsecase {
	return &RefreshUsecase{
		authservice: authservice,
	}
}

func (uc *RefreshUsecase) Execute(refreshToken string) (string, error) {
	return uc.authservice.Refresh(refreshToken)
}
