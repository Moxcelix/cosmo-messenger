package userservice

type UserRepository interface {
	GetUser(userId string) (*User, error)
	CreateUser(user *User) error
	DeleteUser(userId string) error
	UpdateUser(user *User) error
}
