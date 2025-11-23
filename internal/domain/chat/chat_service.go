package chat_domain

type ChatService struct {
	chatRepo ChatRepository
	factory  *ChatFactory
}

func NewChatService(
	chatRepo ChatRepository,
	factory *ChatFactory,
) *ChatService {
	return &ChatService{
		factory:  factory,
		chatRepo: chatRepo,
	}
}

func (s *ChatService) GetDirectChat(user1Id, user2Id string) (*Chat, error) {
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

func (s *ChatService) GetChatById(chatId string) (*Chat, error) {
	chat, err := s.chatRepo.GetByID(chatId)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		return nil, ErrChatNotFound
	}

	return chat, nil
}
