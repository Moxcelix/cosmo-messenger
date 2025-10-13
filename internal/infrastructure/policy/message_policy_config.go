package policy_infrastructure

import (
	"main/internal/config"
	"main/internal/domain/policy"
	"time"
)

type MessagePolicyConfig struct {
	editDuration   time.Duration
	deleteDuration time.Duration
	maxLength      int
	minLength      int
}

func NewMessagePolicyConfig(cfg *config.Config) policy_domain.MessagePolicyConfig {
	return &MessagePolicyConfig{
		editDuration:   cfg.Policies.Message.EditDuration,
		deleteDuration: cfg.Policies.Message.DeleteDuration,
		maxLength:      cfg.Policies.Message.MaxLength,
		minLength:      cfg.Policies.Message.MinLength,
	}
}

func (c *MessagePolicyConfig) EditDuration() time.Duration   { return c.editDuration }
func (c *MessagePolicyConfig) DeleteDuration() time.Duration { return c.deleteDuration }
func (c *MessagePolicyConfig) MaxLength() int                { return c.maxLength }
func (c *MessagePolicyConfig) MinLength() int                { return c.minLength }
