package userservice_application

import (
	userservice "main/internal/domain/userservice"
)

type UserInfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	ID       string `json:"id"`
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
		return nil, userservice.ErrUserNotFound
	}

	return &UserInfo{
		Name:     user.Name,
		Username: user.Username,
		Bio:      user.Bio,
		ID:       user.ID,
	}, nil
}
