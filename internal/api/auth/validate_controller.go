package auth_api

import (
	"github.com/gin-gonic/gin"
	"main/internal/application/auth"
	"net/http"
)

type ValidateController struct {
	usecase *auth_application.ValidateUsecase
}

func NewValidateController(usecase *auth_application.ValidateUsecase) *ValidateController {
	return &ValidateController{
		usecase: usecase,
	}
}

type validateRequest struct {
	AccessToken string `json:"access_token"`
}

type validateResponse struct {
	Username string `json:"username"`
}

// Refresh godoc
// @Summary Validates users access token
// @Description Returns userID if access token is valid
// @Tags auth
// @Accept json
// @Produce json
// @Param input body validateRequest true "Access credentials"
// @Success 200 {object} validateResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/v1/auth/validate [post]
func (c *ValidateController) Validate(ctx *gin.Context) {
	var req validateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username, err := c.usecase.Execute(req.AccessToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, validateResponse{
		Username: username,
	})
}
