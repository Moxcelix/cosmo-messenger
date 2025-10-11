package authservice_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main/internal/application/authservice"
)

type LoginController struct {
	usecase *authservice_application.LoginUsecase
}

func NewLoginController(usecase *authservice_application.LoginUsecase) *LoginController {
	return &LoginController{
		usecase: usecase,
	}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Login godoc
// @Summary User login
// @Description Authenticates user and returns JWT tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param input body loginRequest true "User credentials"
// @Success 200 {object} loginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/v1/auth/login [post]
func (c *LoginController) Login(ctx *gin.Context) {
	var req loginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := c.usecase.Execute(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
