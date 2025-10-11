package user_application

import (
	"main/internal/domain/user"
)

type UserInfo struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	ID       string `json:"id"`
}

type GetInfoUseCase struct {
	repository user_domain.UserRepository
}

func NewGetInfoUseCase(repository user_domain.UserRepository) *GetInfoUseCase {
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
		return nil, user_domain.ErrUserNotFound
	}

	return &UserInfo{
		Name:     user.Name,
		Username: user.Username,
		Bio:      user.Bio,
		ID:       user.ID,
	}, nil
}
