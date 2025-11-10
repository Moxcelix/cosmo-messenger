package user_api

import (
	"main/pkg"
	"net/http"

	user_application "main/internal/application/user"

	"github.com/gin-gonic/gin"
)

type UserDeleteController struct {
	deleteUserUsecase *user_application.DeleteUserUsecase
	logger            pkg.Logger
}

func NewUserDeleteController(
	deleteUserUsecase *user_application.DeleteUserUsecase,
	logger pkg.Logger) *UserDeleteController {
	return &UserDeleteController{
		deleteUserUsecase: deleteUserUsecase,
		logger:            logger,
	}
}

type deleteResponse struct {
	Message string `json:"message"`
}

// DeleteByPath  godoc
// @Summary      Delete user by id (Admin only)
// @Description  Deletes a user by id. Admin access required.
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user_id  path  string  true  "Username of the user to delete"
// @Success      200   {object}  deleteResponse
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Failure      403   {object}  map[string]string
// @Security     BearerAuth
// @Router       /api/v1/users/delete/{user_id} [delete]
func (c *UserDeleteController) DeleteByPath(ctx *gin.Context) {
	userId := ctx.Param("user_id")

	if err := c.deleteUserUsecase.Execute(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, deleteResponse{
		Message: "user deleted successfully",
	})

	c.logger.Infow("user deleted successfully",
		"user_id", userId,
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
	userId := ctx.GetString("UserID")

	if err := c.deleteUserUsecase.Execute(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, deleteResponse{
		Message: "user deleted successfully",
	})

	c.logger.Infow("user deleted successfully",
		"user_id", userId,
	)
}
