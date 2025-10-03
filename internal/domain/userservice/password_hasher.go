package userservice

import (
	"errors"
	"fmt"
	"main/pkg"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrHashingFailed   = errors.New("password hashing failed")
)

type PasswordHasher struct {
	hasher *pkg.Hasher
}

func NewPasswordHasher(hasher *pkg.Hasher) *PasswordHasher {
	return &PasswordHasher{
		hasher: hasher,
	}
}

func (p *PasswordHasher) HashPassword(password string) (string, error) {
	hash, err := p.hasher.Hash([]byte(password))
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (p *PasswordHasher) ValidatePassword(password string, user *User) error {
	computedHash, err := p.hasher.Hash([]byte(password))
	if err != nil {
		return fmt.Errorf("%w: %v", ErrHashingFailed, err)
	}

	if string(computedHash) != user.PasswordHash {
		return ErrInvalidPassword
	}

	return nil
}
