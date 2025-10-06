package userservice

type DeleteUserPolicy interface {
	CanDelete(requester *User, target *User) bool
}

