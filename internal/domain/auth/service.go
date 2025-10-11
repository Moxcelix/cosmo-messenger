package auth_domain

type AuthService interface {
	Login(username, password string) (accessToken, refreshToken string, err error)

	Refresh(refreshToken string) (newAccessToken string, err error)

	ValidateAccessToken(accessToken string) (userID string, err error)
}
