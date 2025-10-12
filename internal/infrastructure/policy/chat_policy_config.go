package policy_infrastructure

import (
	"main/internal/config"
	"main/internal/domain/policy"
)

type ChatPolicyConfig struct {
	maxGroupMembers   int
	maxChatNameLength int
	minChatNameLength int
}

func NewChatPolicyConfig(cfg *config.Config) policy_domain.ChatPolicyConfig {
	return &ChatPolicyConfig{
		maxGroupMembers:   cfg.Policies.ChatPolicyConfig.MaxGroupMembers,
		maxChatNameLength: cfg.Policies.ChatPolicyConfig.MaxChatNameLength,
		minChatNameLength: cfg.Policies.ChatPolicyConfig.MinChatNameLength,
	}
}

func (c *ChatPolicyConfig) MaxGroupMembers() int   { return c.maxGroupMembers }
func (c *ChatPolicyConfig) MaxChatNameLength() int { return c.maxChatNameLength }
func (c *ChatPolicyConfig) MinChatNameLength() int { return c.minChatNameLength }
