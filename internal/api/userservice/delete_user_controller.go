package userservice_api

import (
	"main/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"main/internal/application/userservice"
)

type UserDeleteController struct {
	deleteUserUsecase *userservice_application.DeleteUserUsecase
	logger            pkg.Logger
}

func NewUserDeleteController(
	deleteUserUsecase *userservice_application.DeleteUserUsecase,
	logger pkg.Logger) *UserDeleteController {
	return &UserDeleteController{
		deleteUserUsecase: deleteUserUsecase,
		logger:            logger,
	}
}

type deleteRequest struct {
	Username string `json:"username" binding:"required"`
}

type deleteResponse struct {
	Message string `json:"message"`
}

// @Summary      Delete user
// @Description  Deletes a user by username.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body  body      deleteRequest  true  "Username of the user to delete"
// @Success      200   {object}  deleteResponse
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      403   {object}  map[string]string
// @Router       /api/v1/users/delete [delete]
func (c *UserDeleteController) Delete(ctx *gin.Context) {
	var req deleteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestingUsername := ctx.GetString("Username")
	targetUsername := req.Username

	if err := c.deleteUserUsecase.Execute(requestingUsername, targetUsername); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, deleteResponse{
		Message: "user deleted successfully",
	})

	c.logger.Infow("user deleted successfully",
		"username", req.Username,
	)
}
