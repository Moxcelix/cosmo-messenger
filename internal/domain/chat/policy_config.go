package chat_domain

type ChatPolicyConfig interface {
	MaxGroupMembers() int
	MaxChatNameLength() int
	MinChatNameLength() int
}
