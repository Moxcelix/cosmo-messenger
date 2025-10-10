package userservice_infrastructure

import (
	userservice "main/internal/domain/userservice"
)

const (
	defaultPage  = 1
	defaultCount = 10
	maxPageSize  = 100
)

type GetUsernamesListUsecase struct {
	repository userservice.UserRepository
}

func NewGetUsernamesListUsecase(repository userservice.UserRepository) *GetUsernamesListUsecase {
	return &GetUsernamesListUsecase{
		repository: repository,
	}
}

func (uc *GetUsernamesListUsecase) Execute(page, count int) (*userservice.UsersList, error) {
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
