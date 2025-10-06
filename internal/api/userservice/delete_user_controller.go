package userservice_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main/internal/application/userservice"
)

type DeleteUserController struct {
	deleteUserUsecase *userservice_application.DeleteUserUsecase
}

func NewDeleteUserController(deleteUserUsecase *userservice_application.DeleteUserUsecase) *DeleteUserController {
	return &DeleteUserController{
		deleteUserUsecase: deleteUserUsecase,
	}
}

func (c *DeleteUserController) Delete(ctx gin.Context) {
	
}
