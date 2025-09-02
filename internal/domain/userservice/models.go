package userservice

import (
	"time"
)

type User struct {
	ID           string
	Name         string
	PasswordHash string
	Bio          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
