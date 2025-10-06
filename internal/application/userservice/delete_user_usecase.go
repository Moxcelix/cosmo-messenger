package userservice_application

import (
	userservice "main/internal/domain/userservice"
)

type DeleteUserUsecase struct {
	repository       userservice.UserRepository
	deleteUserPolicy userservice.DeleteUserPolicy
}

func NewDeleteUserUsecase(
	repository userservice.UserRepository,
	deleteUserPolicy userservice.DeleteUserPolicy) *DeleteUserUsecase {
	return &DeleteUserUsecase{
		repository:       repository,
		deleteUserPolicy: deleteUserPolicy,
	}
}

func (uc *DeleteUserUsecase) Execute(requestingUsername, targetUsername string) error {
	targetUser, err := uc.repository.GetUserByUsername(targetUsername)
	if err != nil {
		return err
	}

	if targetUser == nil {
		return ErrUserNotFound
	}

	requestingUser, err := uc.repository.GetUserByUsername(requestingUsername)
	if err != nil {
		return err
	}

	if requestingUser == nil {
		return ErrUserNotFound
	}

	if !uc.deleteUserPolicy.Resolve(requestingUser, targetUser) {
		return ErrNoPermission
	}

	if err := uc.repository.DeleteUser(requestingUsername); err != nil {
		return err
	}

	return nil
}
