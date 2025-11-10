package user_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_application "main/internal/application/user"
)

type UserGetInfoController struct {
	getInfoUseCase *user_application.GetInfoUseCase
}

func NewUserGetInfoController(getInfoUseCase *user_application.GetInfoUseCase) *UserGetInfoController {
	return &UserGetInfoController{
		getInfoUseCase: getInfoUseCase,
	}
}

type infoResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	ID       string `json:"id"`
}

// Get info godoc
// @Summary Get info about user
// @Description Returns information about a user by username
// @Tags users
// @Produce json
// @Param username query string true "Username"
// @Success 200 {object} infoResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/users/get_info [get]
func (c *UserGetInfoController) GetInfo(ctx *gin.Context) {
	username := ctx.Query("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	info, err := c.getInfoUseCase.Execute(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, infoResponse{
		Name:     info.Name,
		Username: info.Username,
		Bio:      info.Bio,
		ID:       info.ID,
	})
}
