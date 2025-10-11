package authservice_api

import (
	"github.com/gin-gonic/gin"
	"main/internal/application/authservice"
	"net/http"
)

type RefreshController struct {
	usecase *authservice_application.RefreshUsecase
}

func NewRefreshController(usecase *authservice_application.RefreshUsecase) *RefreshController {
	return &RefreshController{
		usecase: usecase,
	}
}

type refreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type refreshResponse struct {
	AccessToken string `json:"access_token"`
}

// Refresh godoc
// @Summary User refresh token
// @Description Returns new access token by refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body refreshRequest true "Refresh credentials"
// @Success 200 {object} refreshResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/v1/auth/refresh [post]
func (c *RefreshController) Refresh(ctx *gin.Context) {
	var req refreshRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := c.usecase.Execute(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, refreshResponse{
		AccessToken: accessToken,
	})
}
