package chat_domain

type DirectChatService struct {
	chatRepo    ChatRepository
	chatFactory *ChatFactory
}

func NewDirectChatService(chatRepo ChatRepository, chatFactory *ChatFactory) *DirectChatService {
	return &DirectChatService{
		chatRepo:    chatRepo,
		chatFactory: chatFactory,
	}
}

func (s *DirectChatService) GetDirectChat(user1Id, user2Id string) (*Chat, error) {
	chat, err := s.chatRepo.GetDirectChat(user1Id, user2Id)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		chat, err = s.chatFactory.CreateDirectChat(user1Id, user2Id)
		if err != nil {
			return nil, err
		}
	}

	return chat, err
}
