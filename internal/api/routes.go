package api

import (
	auth_api "main/internal/api/auth"
	message_api "main/internal/api/message"
	ping_api "main/internal/api/ping"
	swagger_api "main/internal/api/swagger"
	user_api "main/internal/api/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	swaggerRoutes *swagger_api.SwaggerRoutes,
	pingRoutes *ping_api.PingRoutes,
	userRoutes *user_api.UserServiceRoutes,
	authRoutes *auth_api.AuthServiceRoutes,
	msgRoutes *message_api.MessageRoutes,
) Routes {
	return Routes{
		swaggerRoutes,
		pingRoutes,
		userRoutes,
		authRoutes,
		msgRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
