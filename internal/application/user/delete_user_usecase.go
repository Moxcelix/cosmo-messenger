package user_application

import (
	"main/internal/domain/user"
)

type DeleteUserUsecase struct {
	repository user_domain.UserRepository
}

func NewDeleteUserUsecase(
	repository user_domain.UserRepository) *DeleteUserUsecase {
	return &DeleteUserUsecase{
		repository: repository,
	}
}

func (uc *DeleteUserUsecase) Execute(username string) error {
	user, err := uc.repository.GetUserByUsername(username)
	if err != nil {
		return err
	}

	if user == nil {
		return user_domain.ErrUserNotFound
	}

	if err := uc.repository.DeleteUserByUsername(username); err != nil {
		return err
	}

	return nil
}
