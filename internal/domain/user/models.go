package user_domain

import "time"

type User struct {
	ID           string
	Name         string
	Username     string
	PasswordHash string
	Bio          string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UsersList struct {
	Users  []*User
	Total  int
	Offset int
	Limit  int
}
