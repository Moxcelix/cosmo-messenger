package auth_infrastructure

import (
	"errors"
	"main/internal/config"
	"main/internal/domain/auth"
	"main/internal/domain/user"
	"main/pkg"
	"time"
)

type InternalAuthService struct {
	jwt            *pkg.Jwt
	accessTTL      time.Duration
	refreshTTL     time.Duration
	userRepo       user_domain.UserRepository
	passwordHasher *user_domain.PasswordHasher
}

type User struct {
	ID       string
	Username string
	Password string
}

func NewInternalAuthService(
	jwt *pkg.Jwt,
	env config.Env,
	userRepo user_domain.UserRepository, 
	passwordHasher *user_domain.PasswordHasher) auth_domain.AuthService {
	return &InternalAuthService{
		jwt:            jwt,
		accessTTL:      env.JwtAccessTTL,
		refreshTTL:     env.JwtRefreshTTL,
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}

func (a *InternalAuthService) Login(username, password string) (string, string, error) {
	user, err := a.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	if err := a.passwordHasher.ValidatePassword(password, user); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := a.jwt.GenerateToken(user.ID, a.accessTTL)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := a.jwt.GenerateToken(user.ID, a.refreshTTL)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *InternalAuthService) Refresh(refreshToken string) (string, error) {
	userID, err := a.jwt.ValidateToken(refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	newAccessToken, err := a.jwt.GenerateToken(userID, a.accessTTL)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}

func (a *InternalAuthService) ValidateAccessToken(accessToken string) (string, error) {
	return a.jwt.ValidateToken(accessToken)
}
