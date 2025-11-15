package broadcasters

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewWebsocketTypingPublisher),
	fx.Provide(NewWebsocketChatPublisher),
)
