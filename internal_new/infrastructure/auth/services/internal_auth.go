package services

import (
	"main/internal/config"
	"main/internal_new/domain/auth/errors"
	"main/internal_new/domain/auth/models"
	"main/internal_new/domain/auth/repositories"
	"main/internal_new/domain/auth/services"
	"main/pkg"
	"time"
)

type InternalAuthService struct {
	jwt            *pkg.Jwt
	accessTTL      time.Duration
	refreshTTL     time.Duration
	userRepo       repositories.UserRepository
	passwordHasher *PasswordHasher
}

func NewInternalAuthService(
	jwt *pkg.Jwt,
	env config.Env,
	userRepo repositories.UserRepository,
	passwordHasher *PasswordHasher,
) services.AuthService {
	return &InternalAuthService{
		jwt:            jwt,
		accessTTL:      env.JwtAccessTTL,
		refreshTTL:     env.JwtRefreshTTL,
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}

func (a *InternalAuthService) Login(user *models.User, password string) (*models.Login, error) {
	if err := a.passwordHasher.ValidatePassword(password, user); err != nil {
		return nil, err
	}

	accessToken, err := a.jwt.GenerateToken(user.ID, a.accessTTL)
	if err != nil {
		return nil, err
	}

	refreshToken, err := a.jwt.GenerateToken(user.ID, a.refreshTTL)
	if err != nil {
		return nil, err
	}

	return &models.Login{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *InternalAuthService) Refresh(refreshToken string) (*models.Refresh, error) {
	userID, err := a.jwt.ValidateToken(refreshToken)
	if err != nil {
		return nil, errors.ErrInvalidRefreshToken
	}

	newAccessToken, err := a.jwt.GenerateToken(userID, a.accessTTL)
	if err != nil {
		return nil, err
	}

	return &models.Refresh{
		AccessToken: newAccessToken,
	}, nil
}

func (a *InternalAuthService) ValidateAccessToken(accessToken string) (*models.Validate, error) {
	userId, err := a.jwt.ValidateToken(accessToken)
	if err != nil {
		return nil, err
	}
	return &models.Validate{
		UserID: userId,
	}, nil
}

func (a *InternalAuthService) Register(name string, username string, password string, bio string) error {
	existing, err := a.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.ErrUsernameAlreadyTaken
	}

	hash, err := a.passwordHasher.HashPassword(password)
	if err != nil {
		return err
	}

	now := time.Now()
	user := &models.User{
		Name:         name,
		Username:     username,
		PasswordHash: string(hash),
		Bio:          bio,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	return a.userRepo.CreateUser(user)
}
