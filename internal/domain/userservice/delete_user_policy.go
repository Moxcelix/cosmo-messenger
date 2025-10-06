package userservice

type DeleteUserPolicy interface {
	Resolve(requester *User, target *User) bool
}

type DefaultDeleteUserPolicy struct{}

func NewDefaultDeleteUserPolicy() DeleteUserPolicy {
	return &DefaultDeleteUserPolicy{}
}

func (p *DefaultDeleteUserPolicy) Resolve(requester *User, target *User) bool {
	return requester.ID == target.ID
}
