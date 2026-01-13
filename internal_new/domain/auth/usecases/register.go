package usecases

import "main/internal_new/domain/auth/services"

type RegisterUseCase struct {
	authservice services.AuthService
}

func NewRegisterUseCase(authservice services.AuthService) *RegisterUseCase {
	return &RegisterUseCase{
		authservice: authservice,
	}
}

func (r *RegisterUseCase) Execute(name string, username string, password string, bio string) error {
	return r.authservice.Register(name, username, password, bio)
}
