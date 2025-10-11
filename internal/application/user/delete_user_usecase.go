package userservice_application

import (
	userservice "main/internal/domain/userservice"
)

type DeleteUserUsecase struct {
	repository userservice.UserRepository
}

func NewDeleteUserUsecase(
	repository userservice.UserRepository) *DeleteUserUsecase {
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
		return userservice.ErrUserNotFound
	}

	if err := uc.repository.DeleteUserByUsername(username); err != nil {
		return err
	}

	return nil
}
