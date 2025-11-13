package controllers

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewDirectMessageController),
	fx.Provide(NewGetHistoryController),
	fx.Provide(NewGetDirectMessagesController),
	fx.Provide(NewSendMessageController),
)
