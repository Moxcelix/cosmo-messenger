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

type deleteResponse struct {
	Message string `json:"message"`
}

// @Summary      Delete user by username (Admin only)
// @Description  Deletes a user by username. Admin access required.
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        username  path  string  true  "Username of the user to delete"
// @Success      200   {object}  deleteResponse
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      403   {object}  map[string]string
// @Security     BearerAuth
// @Router       /api/v1/users/admin/delete/{username} [delete]
func (c *UserDeleteController) DeleteByPath(ctx *gin.Context) {
	username := ctx.Param("Username")

	if err := c.deleteUserUsecase.Execute(username); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, deleteResponse{
		Message: "user deleted successfully",
	})

	c.logger.Infow("user deleted successfully",
		"username", username,
	)
}

// @Summary      Delete current user
// @Description  Deletes the currently authenticated user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200   {object}  deleteResponse
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      403   {object}  map[string]string
// @Security     BearerAuth
// @Router       /api/v1/users/delete [delete]
func (c *UserDeleteController) DeleteByContext(ctx *gin.Context) {
	username := ctx.GetString("Username")

	if err := c.deleteUserUsecase.Execute(username); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, deleteResponse{
		Message: "user deleted successfully",
	})

	c.logger.Infow("user deleted successfully",
		"username", username,
	)
}
