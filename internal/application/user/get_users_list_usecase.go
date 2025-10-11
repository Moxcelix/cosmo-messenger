package userservice_application

import (
	userservice "main/internal/domain/userservice"
)

const (
	defaultPage  = 1
	defaultCount = 10
	maxPageSize  = 100
)

type GetUsersListUsecase struct {
	repository userservice.UserRepository
}

func NewGetUsersListUsecase(repository userservice.UserRepository) *GetUsersListUsecase {
	return &GetUsersListUsecase{
		repository: repository,
	}
}

func (uc *GetUsersListUsecase) Execute(page, count int) (*userservice.UsersList, error) {
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
