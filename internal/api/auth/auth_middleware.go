package auth_api

import (
	"main/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	authapp "main/internal/application/auth"
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

		token, err := pkg.ParseBearerToken(authHeader)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

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
