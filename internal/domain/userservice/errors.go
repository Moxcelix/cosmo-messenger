package userservice

import (
	"errors"
)

var ErrUserNotFound = errors.New("user not found")
var ErrNoPermission = errors.New("user does not have permission")
var ErrUsernameAlreadyTaken = errors.New("username already taken")
var ErrRequestingUserNotFound = errors.New("requesting user not found")
var ErrTargetUserNotFound = errors.New("target user not found")
