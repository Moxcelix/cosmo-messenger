package repositories

import "main/internal_new/domain/auth/models"

type UserRepository interface {
	GetUserById(userId string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
	DeleteUserByUsername(username string) error
	DeleteUserById(userId string) error
	UpdateUser(user *models.User) error
	UserExistsById(userId string) (bool, error)
	UserExistsByUsername(username string) (bool, error)
}
