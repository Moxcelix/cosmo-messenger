package chat_domain

type ChatService struct {
	factory    *ChatFactory
	chatRepo   ChatRepository
	chatPolicy *ChatPolicy
}

func NewChatService(
	factory *ChatFactory,
	chatRepo ChatRepository,
	chatPolicy *ChatPolicy,
) *ChatService {
	return &ChatService{
		factory:    factory,
		chatRepo:   chatRepo,
		chatPolicy: chatPolicy,
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

func (s *ChatService) GetChatForUser(chatId, userId string) (*Chat, error) {
	chat, err := s.chatRepo.GetByID(chatId)
	if err != nil {
		return nil, err
	}

	if chat == nil {
		return nil, ErrChatNotFound
	}

	if err := s.chatPolicy.ValidateUserAccess(userId, chat); err != nil {
		return nil, err
	}

	return chat, nil
}
