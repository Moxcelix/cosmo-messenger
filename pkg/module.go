package pkg

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(GetLogger),
	fx.Provide(NewHasher),
	fx.Provide(NewMongoDatabase),
	fx.Provide(NewJwt),
	fx.Provide(NewPostgresDatabase),
)
