package chat_domain

import (
	"errors"
)

var (
	ErrShortChatName   = errors.New("chat name is to short")
	ErrLongChatName    = errors.New("chat name is to long")
	ErrManyChatMembers = errors.New("too many members in chat")
)

type ChatPolicy struct {
	cfg ChatPolicyConfig
}

func NewChatPolicy(cfg ChatPolicyConfig) *ChatPolicy {
	return &ChatPolicy{
		cfg: cfg,
	}
}

func (p *ChatPolicy) ValidateChatName(chatName string) error {
	if len(chatName) < p.cfg.MinChatNameLength() {
		return ErrShortChatName
	}

	if len(chatName) > p.cfg.MaxChatNameLength() {
		return ErrLongChatName
	}

	return nil
}

func (p *ChatPolicy) ValidateGroupMembers(count int) error {
	if count > p.cfg.MaxGroupMembers() {
		return ErrManyChatMembers
	}

	return nil
}

func (p *ChatPolicy) ValidateUserAccess(userId string, chat *Chat) error {
	hasAccess := chat.HasMember(userId)
	if !hasAccess {
		return ErrChatAccessDenied
	}
	return nil
}
