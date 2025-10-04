package authservice_api

import (
	"github.com/gin-gonic/gin"
	"main/internal/application/authservice"
	"net/http"
)

type ValidateController struct {
	usecase *authservice_application.ValidateUsecase
}

func NewValidateController(usecase *authservice_application.ValidateUsecase) *ValidateController {
	return &ValidateController{
		usecase: usecase,
	}
}

type validateRequest struct {
	accessToken string `json:"accsess_token"`
}

type validateResponse struct {
	userID string `json:"user_id"`
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

	userID, err := c.usecase.Execute(req.accessToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, validateResponse{
		userID: userID,
	})
}
