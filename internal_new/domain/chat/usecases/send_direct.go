package usecases

type SendDirectMessageUsecase struct {
}

func NewSendDirectMessageUsecase() *SendDirectMessageUsecase {
	return &SendDirectMessageUsecase{}
}

func (uc *SendDirectMessageUsecase) Execute() {}
