package userservice_application

import (
	userservice "main/internal/domain/userservice"
)

const (
	defaultPage  = 1
	defaultCount = 10
	maxPageSize  = 100
)

type usernamesList struct {
	Usernames []string
	Total     int
	Offset    int
	Limit     int
}

type GetUsernamesListUsecase struct {
	repository userservice.UserRepository
}

func NewGetUsernamesListUsecase(repository userservice.UserRepository) *GetUsernamesListUsecase {
	return &GetUsernamesListUsecase{
		repository: repository,
	}
}

func (uc *GetUsernamesListUsecase) Execute(page, count int) (*usernamesList, error) {
	if page < 1 {
		page = defaultPage
	}
	if count < 1 {
		count = defaultCount
	}
	if count > maxPageSize {
		count = maxPageSize
	}

	userList, err := uc.repository.GetUsersByRange((page-1)*count, count)

	if err != nil {
		return nil, err
	}
	usernames := make([]string, len(userList.Users))
	for i, user := range userList.Users {
		usernames[i] = user.Username
	}

	result := &usernamesList{
		Usernames: usernames,
		Total:     userList.Total,
		Limit:     userList.Limit,
		Offset:    userList.Offset,
	}

	return result, nil
}
