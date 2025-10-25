package user_application

import (
	"time"

	user "main/internal/domain/user"
	"main/pkg"
)

type RegisterUseCase struct {
	repository user.UserRepository
	hasher     *pkg.Hasher
}

func NewRegisterUseCase(repository user.UserRepository, hasher *pkg.Hasher) *RegisterUseCase {
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
		return user.ErrUsernameAlreadyTaken
	}

	hash, err := r.hasher.Hash([]byte(password))
	if err != nil {
		return err
	}

	now := time.Now()
	user := &user.User{
		Name:         name,
		Username:     username,
		PasswordHash: string(hash),
		Bio:          bio,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	return r.repository.CreateUser(user)
}
