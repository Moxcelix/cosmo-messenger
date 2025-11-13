package controllers

import (
	"net/http"

	user_application "main/internal/application/user/usecases"
	"main/pkg"

	"github.com/gin-gonic/gin"
)

type UserRegisterController struct {
	registerUseCase *user_application.RegisterUseCase
	logger          pkg.Logger
}

func NewUserRegisterController(registerUseCase *user_application.RegisterUseCase, logger pkg.Logger) *UserRegisterController {
	return &UserRegisterController{
		registerUseCase: registerUseCase,
		logger:          logger,
	}
}

type registerRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Bio      string `json:"bio"`
}

type registerResponse struct {
	Message string `json:"message"`
}

// Register godoc
// @Summary User registration
// @Description Creates a new user in system
// @Tags users
// @Accept json
// @Produce json
// @Param input body registerRequest true "User data"
// @Success 201 {object} registerResponse
// @Failure 400 {object} map[string]string
// @Router /api/v1/users/register [post]
func (c *UserRegisterController) Register(ctx *gin.Context) {
	var req registerRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.registerUseCase.Execute(req.Name, req.Username, req.Password, req.Bio)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, registerResponse{
		Message: "user registered successfully",
	})

	c.logger.Infow("user registered successfully",
		"name", req.Name,
		"username", req.Username,
	)
}
