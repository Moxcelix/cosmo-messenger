package usecases

import (
	user_domain "main/internal/domain/user"
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

func (uc *DeleteUserUsecase) Execute(userId string) error {
	exists, err := uc.repository.UserExists(userId)
	if err != nil {
		return err
	}

	if !exists {
		return user_domain.ErrUserNotFound
	}

	if err := uc.repository.DeleteUserById(userId); err != nil {
		return err
	}

	return nil
}
