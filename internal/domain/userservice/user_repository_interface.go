package userservice

type UserRepository interface {
	GetUser(userId int64) (*User, error)
	CreateUser(user *User) error
	DeleteUser(userId int64) error
	UpdateUser(user *User) error
}
