package usecases

import (
	"time"

	errors "main/internal_new/domain/auth/errors"
	"main/internal_new/domain/auth/models"
	"main/internal_new/domain/auth/repositories"
	"main/pkg"
)

type RegisterUseCase struct {
	repository repositories.UserRepository
	hasher     *pkg.Hasher
}

func NewRegisterUseCase(
	repository repositories.UserRepository,
	hasher *pkg.Hasher) *RegisterUseCase {
	return &RegisterUseCase{
		repository: repository,
		hasher:     hasher,
	}
}

func (r *RegisterUseCase) Execute(name string, username string, password string, bio string) error {
	existing, err := r.repository.GetUserByUsername(username)
	if err != nil {
		return err
	}
	if existing != nil {
		return errors.ErrUsernameAlreadyTaken
	}

	hash, err := r.hasher.Hash([]byte(password))
	if err != nil {
		return err
	}
	// TODO: transport to service
	now := time.Now()
	user := &models.User{
		Name:         name,
		Username:     username,
		PasswordHash: string(hash),
		Bio:          bio,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	return r.repository.CreateUser(user)
}
