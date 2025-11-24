package auth_domain

import user_domain "main/internal/domain/user"

type AuthService interface {
	Login(user *user_domain.User, password string) (accessToken, refreshToken string, err error)

	Refresh(refreshToken string) (newAccessToken string, err error)

	ValidateAccessToken(accessToken string) (userID string, err error)
}
