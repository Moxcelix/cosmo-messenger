package usecases

type SendChatMessageUsecase struct {
}

func NewSendChatMessageUsecase() *SendChatMessageUsecase {
	return &SendChatMessageUsecase{}
}

func (uc *SendChatMessageUsecase) Execute() {}
