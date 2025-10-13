package bootstrap

import (
	"main/internal/api"
	"main/internal/api/auth"
	ping_api "main/internal/api/ping"
	swagger_api "main/internal/api/swagger"
	user_api "main/internal/api/user"
	"main/internal/application/auth"
	user_application "main/internal/application/user"
	"main/internal/config"
	"main/internal/domain/chat"
	"main/internal/domain/message"
	user "main/internal/domain/user"
	"main/internal/infrastructure/auth"
	"main/internal/infrastructure/chat"
	"main/internal/infrastructure/message"
	user_infrastructure "main/internal/infrastructure/user"
	"main/pkg"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	api.Module,
	config.Module,
	pkg.Module,
	ping_api.Module,
	user_infrastructure.Module,
	user_application.Module,
	swagger_api.Module,
	user_api.Module,
	user.Module,
	auth_infrastructure.Module,
	auth_application.Module,
	auth_api.Module,
	message_domain.Module,
	message_infrastructure.Module,
	chat_domain.Module,
	chat_infrastructure.Module,
)
