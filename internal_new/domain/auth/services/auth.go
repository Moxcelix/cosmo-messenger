package services

import "main/internal_new/domain/auth/models"

// TODO: make register func, move hashing to infra

type AuthService interface {
	Register(name string, username string, password string, bio string) error

	Login(user *models.User, password string) (*models.Login, error)

	Refresh(refreshToken string) (*models.Refresh, error)

	ValidateAccessToken(accessToken string) (*models.Validate, error)
}
