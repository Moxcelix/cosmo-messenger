package middlewares

import (
	"errors"
	"main/internal/config"
	"main/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminAuthMiddleware struct {
	adminToken string
}

func NewAdminAuthMiddleware(env config.Env) *AdminAuthMiddleware {
	return &AdminAuthMiddleware{
		adminToken: env.AdminToken,
	}
}

func (m *AdminAuthMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		token, err := pkg.ParseBearerToken(authHeader)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if token != m.adminToken {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": errors.New("access token is not valid"),
			})
			return
		}

		ctx.Next()
	}
}
