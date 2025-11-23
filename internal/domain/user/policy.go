package user_domain

type UserPolicy struct {
	userRepo UserRepository
}

func NewUserPolicy(userRepo UserRepository) *UserPolicy {
	return &UserPolicy{
		userRepo: userRepo,
	}
}

func (p *UserPolicy) ValidateExists(userId string) error {
	exists, err := p.userRepo.UserExists(userId)
	if err != nil {
		return err
	}
	if !exists {
		return ErrUserNotFound
	}
	return nil
}

func (p *UserPolicy) ValidateExistsByUsername(username string) error {
	user, err := p.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}
	if user == nil {
		return ErrUserNotFound
	}
	return nil
}
