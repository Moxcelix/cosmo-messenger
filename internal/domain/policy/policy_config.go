package policy_domain

import "time"

type MessagePolicyConfig interface {
	EditDuration() time.Duration
	DeleteDuration() time.Duration
	MaxLength() int
	MinLength() int
}

type ChatPolicyConfig interface {
	MaxGroupMembers() int
	MaxChatNameLength() int
	MinChatNameLength() int
}
