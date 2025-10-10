package userservice_api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userservice_application "main/internal/application/userservice"
)

type GetUsernamesListController struct {
	getUsernamesListUsecase *userservice_application.GetUsernamesListUsecase
}

func NewGetUsernamesListController(
	getUsernamesListUsecase *userservice_application.GetUsernamesListUsecase) *GetUsernamesListController {
	return &GetUsernamesListController{
		getUsernamesListUsecase: getUsernamesListUsecase,
	}
}

type listResponse struct {
	Usernames []string `json:"usernames"`
	Total     int      `json:"total"`
	Offset    int      `json:"offset"`
	Limit     int      `json:"limit"`
}

// GetUsernamesList godoc
// @Summary Get usernames list
// @Description Returns paginated usernames list
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param count query int false "Items per page" default(10)
// @Success 200 {object} listResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/get_usernames_list [get]
func (c *GetUsernamesListController) GetList(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid page parameter"})
		return
	}

	countStr := ctx.DefaultQuery("count", "10")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid count parameter"})
		return
	}

	list, err := c.getUsernamesListUsecase.Execute(page, count)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, listResponse{
		Usernames: list.Usernames,
		Total:     list.Total,
		Offset:    list.Offset,
		Limit:     list.Limit,
	})
}
