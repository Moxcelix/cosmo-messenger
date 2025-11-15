package application

import (
	auth_application "main/internal/application/auth"
	chat_application "main/internal/application/chat"
	message_application "main/internal/application/message"
	user_application "main/internal/application/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	auth_application.Module,
	chat_application.Module,
	user_application.Module,
	message_application.Module,
)
