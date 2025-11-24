package user_domain

type UniqueService struct {
	userRepo UserRepository
}

func NewUniqueService(userRepo UserRepository) *UniqueService {
	return &UniqueService{
		userRepo: userRepo,
	}
}

func (s *UniqueService) CheckUsernameAvailable(username string) error {
	exists, err := s.userRepo.UserExistsByUsername(username)
	if err != nil {
		return err
	}

	if exists {
		return ErrUsernameAlreadyTaken
	}

	return nil
}
