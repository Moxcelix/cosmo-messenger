package message_domain

import "time"

type MessagePolicyConfig interface {
	EditDuration() time.Duration
	DeleteDuration() time.Duration
	MaxLength() int
	MinLength() int
}
