package policy_domain

import (
	"errors"
	"main/internal/domain/message"
	"time"
)

var (
	ErrEditDurationExpired   = errors.New("edit time expired")
	ErrDeleteDurationExpired = errors.New("delete time expired")
	ErrMinLength             = errors.New("too small message")
	ErrMaxLength             = errors.New("too large message")
)

type MessagePolicy struct {
	cfg MessagePolicyConfig
}

func NewMessagePolicy(cfg MessagePolicyConfig) *MessagePolicy {
	return &MessagePolicy{
		cfg: cfg,
	}
}

func (p *MessagePolicy) ValidateMessageContent(message string) error {
	if len(message) < p.cfg.MinLength() {
		return ErrMinLength
	}

	if len(message) > p.cfg.MaxLength() {
		return ErrMaxLength
	}

	return nil
}

func (p *MessagePolicy) ValidateDelete(msg message_domain.Message, now time.Time) error {
	if now.Sub(msg.CreatedAt) > p.cfg.DeleteDuration() {
		return ErrDeleteDurationExpired
	}
	return nil
}

func (p *MessagePolicy) ValidateEdit(msg message_domain.Message, now time.Time) error {
	if now.Sub(msg.CreatedAt) > p.cfg.EditDuration() {
		return ErrEditDurationExpired
	}
	return nil
}
