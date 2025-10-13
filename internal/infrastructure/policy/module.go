package policy_infrastructure

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewChatPolicyConfig),
	fx.Provide(NewMessagePolicyConfig),
)
