package user_domain

import (
	"errors"
)

var ErrUserNotFound = errors.New("user not found")
var ErrNoPermission = errors.New("user does not have permission")
var ErrUsernameAlreadyTaken = errors.New("username already taken")
var ErrUsersInRangeNotFound = errors.New("users in range not found")
