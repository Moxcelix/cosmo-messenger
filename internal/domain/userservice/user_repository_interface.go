package userservice

type UserRepository interface {
	GetUserById(userId string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	CreateUser(user *User) error
	DeleteUserById(userId string) error
	DeleteUserByUsername(userId string) error
	UpdateUser(user *User) error
	GetUsersByRange(offset, ligimt int) (*UsersList, error)
}
