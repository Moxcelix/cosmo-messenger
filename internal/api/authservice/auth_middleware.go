package authservice_api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	authapp "main/internal/application/authservice"
)

type AuthMiddleware struct {
	validateUsecase *authapp.ValidateUsecase
}

func NewAuthMiddleware(validateUsecase *authapp.ValidateUsecase) *AuthMiddleware {
	return &AuthMiddleware{
		validateUsecase: validateUsecase,
	}
}

func (m *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header format",
			})
			return
		}

		token := parts[1]

		username, err := m.validateUsecase.Execute(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.Set("Username", username)

		ctx.Next()
	}
}
