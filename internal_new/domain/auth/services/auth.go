package services

import "main/internal_new/domain/auth/models"

type AuthService interface {
	Login(user *models.User, password string) (accessToken, refreshToken string, err error)

	Refresh(refreshToken string) (newAccessToken string, err error)

	ValidateAccessToken(accessToken string) (userID string, err error)
}
