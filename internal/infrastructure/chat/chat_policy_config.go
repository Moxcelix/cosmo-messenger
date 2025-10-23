package chat_infrastructure

import (
	"main/internal/config"
	"main/internal/domain/chat"
)

type ChatPolicyConfig struct {
	maxGroupMembers   int
	maxChatNameLength int
	minChatNameLength int
}

func NewChatPolicyConfig(cfg *config.Config) chat_domain.ChatPolicyConfig {
	return &ChatPolicyConfig{
		maxGroupMembers:   cfg.Policies.Chat.MaxGroupMembers,
		maxChatNameLength: cfg.Policies.Chat.MaxChatNameLength,
		minChatNameLength: cfg.Policies.Chat.MinChatNameLength,
	}
}

func (c *ChatPolicyConfig) MaxGroupMembers() int   { return c.maxGroupMembers }
func (c *ChatPolicyConfig) MaxChatNameLength() int { return c.maxChatNameLength }
func (c *ChatPolicyConfig) MinChatNameLength() int { return c.minChatNameLength }
