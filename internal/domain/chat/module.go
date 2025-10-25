package chat_domain

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewChatPolicy),
	fx.Provide(NewDirectChatProvider),
	fx.Provide(NewChatNamingService),
)
