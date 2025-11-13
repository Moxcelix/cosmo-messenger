package middlewares

import (
	"net/http"

	authapp "main/internal/application/auth/usecases"

	"github.com/gin-gonic/gin"
)

type QueryAuthMiddleware struct {
	validateUsecase *authapp.ValidateUsecase
}

func NewQueryAuthMiddleware(validateUsecase *authapp.ValidateUsecase) *QueryAuthMiddleware {
	return &QueryAuthMiddleware{
		validateUsecase: validateUsecase,
	}
}

func (m *QueryAuthMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query("token")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token query parameter required for WebSocket connection",
			})
			return
		}

		userId, err := m.validateUsecase.Execute(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.Set("UserID", userId)
		ctx.Next()
	}
}
