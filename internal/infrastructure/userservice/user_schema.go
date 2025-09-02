package userservice_infrastructure

import (
	"time"
)

type User struct {
	ID           string    `json:"id" bson:"_id"`
	Name         string    `json:"name" bson:"name"`
	PasswordHash string    `json:"password_hash" bson:"password_hash"`
	Bio          string    `json:"bio" bson:"bio"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}
