package usecases

import (
	user_domain "main/internal/domain/user"
	auth_repositories "main/internal_new/domain/auth/repositories"
	"main/internal_new/domain/chat/projections"
	"main/internal_new/domain/chat/queries"
	"main/internal_new/domain/chat/services"
)

type DirectHistoryUsecase struct {
	historyQuery          queries.ChatHistoryQuery
	userRepo              auth_repositories.UserRepository
	directService         services.DirectChatService
	messageDefaultService *services.MessageDefaultService
	chatHeaderService     *services.ChatHeaderService
	messageReplyService   *services.MessageReplyService
	senderService         *services.SenderService
}

func NewDirectHistoryUsecase(
	historyQuery queries.ChatHistoryQuery,
	userRepo auth_repositories.UserRepository,
	directService services.DirectChatService,
	messageDefaultService *services.MessageDefaultService,
	chatHeaderService *services.ChatHeaderService,
	messageReplyService *services.MessageReplyService,
	senderService *services.SenderService,
) *DirectHistoryUsecase {
	return &DirectHistoryUsecase{
		historyQuery:          historyQuery,
		userRepo:              userRepo,
		directService:         directService,
		messageDefaultService: messageDefaultService,
		chatHeaderService:     chatHeaderService,
		messageReplyService:   messageReplyService,
		senderService:         senderService,
	}
}

func (uc *DirectHistoryUsecase) Execute(
	userId, targetUsername, cursorMessageId string, count int, direction string,
) (
	*projections.ChatHeader,
	map[string]*projections.MessageDefault,
	map[string]*projections.MessageReply,
	map[string]*projections.Sender,
	bool,
	bool,
	error,
) {
	// TODO: separated user-service???
	companion, err := uc.userRepo.GetUserByUsername(targetUsername)
	if err != nil {
		return nil, nil, nil, nil, false, false, err
	}

	if companion == nil {
		return nil, nil, nil, nil, false, false, user_domain.ErrUserNotFound
	}

	chat, err := uc.directService.GetDirectChat(userId, companion.ID)
	if err != nil {
		return nil, nil, nil, nil, false, false, err
	}

	messages, replies, users, hasNext, hasPrev, err := uc.historyQuery.Query(chat.ID, cursorMessageId, count, direction)
	if err != nil {
		return nil, nil, nil, nil, false, false, err
	}

	defaultMessages := uc.messageDefaultService.ProjectMessages(messages)
	chatHeader := uc.chatHeaderService.ProjectChat(chat, users, userId)
	replyMessages := uc.messageReplyService.ProjectMessages(replies)
	senders := uc.senderService.ProjectUsers(users)

	return chatHeader, defaultMessages, replyMessages, senders, hasNext, hasPrev, nil
}
