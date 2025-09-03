package userservice

type UserRepository interface {
	GetUserById(userId string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	CreateUser(user *User) error
	DeleteUser(userId string) error
	UpdateUser(user *User) error
}
