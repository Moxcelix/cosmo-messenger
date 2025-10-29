package bootstrap

import (
	"main/internal/api"
	auth_api "main/internal/api/auth"
	chat_api "main/internal/api/chat"
	message_api "main/internal/api/message"
	ping_api "main/internal/api/ping"
	swagger_api "main/internal/api/swagger"
	user_api "main/internal/api/user"
	websocket_api "main/internal/api/websocket"
	auth_application "main/internal/application/auth"
	chat_application "main/internal/application/chat"
	message_application "main/internal/application/message"
	user_application "main/internal/application/user"
	"main/internal/config"
	chat_domain "main/internal/domain/chat"
	message_domain "main/internal/domain/message"
	user_domain "main/internal/domain/user"
	auth_infrastructure "main/internal/infrastructure/auth"
	chat_infrastructure "main/internal/infrastructure/chat"
	message_infrastructure "main/internal/infrastructure/message"
	user_infrastructure "main/internal/infrastructure/user"
	"main/pkg"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	api.Module,
	config.Module,
	pkg.Module,
	ping_api.Module,

	user_domain.Module,
	user_application.Module,
	user_infrastructure.Module,
	user_api.Module,

	swagger_api.Module,

	auth_application.Module,
	auth_infrastructure.Module,
	auth_api.Module,

	message_domain.Module,
	message_application.Module,
	message_infrastructure.Module,
	message_api.Module,

	chat_domain.Module,
	chat_application.Module,
	chat_infrastructure.Module,
	chat_api.Module,

	websocket_api.Module,
)
