package chat_domain

import "errors"

var (
	ErrChatAlreadyExsists       = errors.New("chat already exists")
	ErrChatNotFound             = errors.New("chat not found")
	ErrCannotCreateChatWithSelf = errors.New("cannot create chat with self")
)
