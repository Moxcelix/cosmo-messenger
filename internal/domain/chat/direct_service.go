package chat_domain

type DirectChatService struct {
	factory  *ChatFactory
	chatRepo ChatRepository
}

func NewDirectChatService(
	factory *ChatFactory,
	chatRepo ChatRepository,
) *DirectChatService {
	return &DirectChatService{
		factory:  factory,
		chatRepo: chatRepo,
	}
}

func (s *DirectChatService) GetDirectChat(user1Id, user2Id string) (*Chat, error) {
	chat, err := s.chatRepo.GetDirectChat(user1Id, user2Id)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		chat, err = s.factory.CreateDirectChat(user1Id, user2Id)
		if err != nil {
			return nil, err
		}
	}

	return chat, err
}
