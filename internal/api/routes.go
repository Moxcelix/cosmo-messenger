package api

import (
	authservice_api "main/internal/api/authservice"
	ping_api "main/internal/api/ping"
	swagger_api "main/internal/api/swagger"
	userservice_api "main/internal/api/userservice"

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
	userRoutes *userservice_api.UserServiceRoutes,
	authRoutes *authservice_api.AuthServiceRoutes,
) Routes {
	return Routes{
		swaggerRoutes,
		pingRoutes,
		userRoutes,
		authRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
