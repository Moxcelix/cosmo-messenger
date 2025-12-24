package services

import "main/internal_new/domain/auth/models"

// TODO: make register func, move hashing to infra

type AuthService interface {
	Login(user *models.User, password string) (accessToken, refreshToken string, err error)

	Refresh(refreshToken string) (newAccessToken string, err error)

	ValidateAccessToken(accessToken string) (userID string, err error)
}
