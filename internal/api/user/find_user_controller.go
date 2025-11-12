package user_api

import (
	"errors"
	user_application "main/internal/application/user/usecases"
	user_domain "main/internal/domain/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FindUserController struct {
	usecase *user_application.FindUserUsecase
}

func NewFindUserController(usecase *user_application.FindUserUsecase) *FindUserController {
	return &FindUserController{
		usecase: usecase,
	}
}

// Find user godoc
// @Summary Find user
// @Description Returns user by username for direct chating
// @Tags users
// @Produce json
// @Param username query string true "Username"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/find [get]
func (c *FindUserController) FindUser(ctx *gin.Context) {
	userId := ctx.GetString("UserID")
	targetUsername := ctx.Query("username")
	if targetUsername == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	info, err := c.usecase.Execute(userId, targetUsername)

	if err != nil {
		switch {
		case errors.Is(err, user_domain.ErrUserNotFound):
			ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, info)
}
