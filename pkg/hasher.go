package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct {
}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(bytes []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return hash, nil
}
