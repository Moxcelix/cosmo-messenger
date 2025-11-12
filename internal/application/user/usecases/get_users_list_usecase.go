package usecases

import (
	user "main/internal/domain/user"
)

const (
	defaultPage  = 1
	defaultCount = 10
	maxPageSize  = 100
)

type GetUsersListUsecase struct {
	repository user.UserRepository
}

func NewGetUsersListUsecase(repository user.UserRepository) *GetUsersListUsecase {
	return &GetUsersListUsecase{
		repository: repository,
	}
}

func (uc *GetUsersListUsecase) Execute(page, count int) (*user.UsersList, error) {
	if page < 1 {
		page = defaultPage
	}
	if count < 1 {
		count = defaultCount
	}
	if count > maxPageSize {
		count = maxPageSize
	}

	return uc.repository.GetUsersByRange((page-1)*count, count)
}
