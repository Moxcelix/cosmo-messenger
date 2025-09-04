package userservice_application

import (
	"errors"
	userservice "main/internal/domain/userservice"
)

var ErrUserNotFound = errors.New("user not found")

type UserInfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
}

type GetInfoUseCase struct {
	repository userservice.UserRepository
}

func NewGetInfoUseCase(repository userservice.UserRepository) *GetInfoUseCase {
	return &GetInfoUseCase{
		repository: repository,
	}
}

func (uc *GetInfoUseCase) Execute(username string) (*UserInfo, error) {
	user, err := uc.repository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErrUserNotFound
	}

	return &UserInfo{
		Name:     user.Name,
		Username: user.Username,
		Bio:      user.Bio,
	}, nil
}
