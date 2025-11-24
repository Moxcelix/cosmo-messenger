package user_domain

type UserRepository interface {
	GetUserById(userId string) (*User, error)
	GetUserByUsername(username string) (*User, error)
	CreateUser(user *User) error
	DeleteUserByUsername(username string) error
	DeleteUserById(userId string) error
	UpdateUser(user *User) error
	GetUsersByRange(offset, limit int) (*UsersList, error)
	UserExistsById(userId string) (bool, error)
	UserExistsByUsername(username string) (bool, error)
}
