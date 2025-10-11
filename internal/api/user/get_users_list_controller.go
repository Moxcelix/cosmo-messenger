package user_api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	user_application "main/internal/application/user"
)

type GetUsersListController struct {
	getUsersListUsecase *user_application.GetUsersListUsecase
}

func NewGetUsersListController(
	getUsersListUsecase *user_application.GetUsersListUsecase) *GetUsersListController {
	return &GetUsersListController{
		getUsersListUsecase: getUsersListUsecase,
	}
}

type usernamesListResponse struct {
	Usernames []string       `json:"usernames"`
	Meta      paginationMeta `json:"meta"`
}

type paginationMeta struct {
	HasPrev    bool `json:"has_prev"`
	HasNext    bool `json:"has_next"`
	Page       int  `json:"page"`
	TotalPages int  `json:"total_pages"`
	Total      int  `json:"total"`
}

// GetUsernamesList godoc
// @Summary Get usernames list
// @Description Returns paginated usernames list
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param count query int false "Items per page" default(10)
// @Success 200 {object} usernamesListResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/users/get_usernames_list [get]
func (c *GetUsersListController) GetUsernameList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid page parameter"})
		return
	}

	count, err := strconv.Atoi(ctx.DefaultQuery("count", "10"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid count parameter"})
		return
	}

	list, err := c.getUsersListUsecase.Execute(page, count)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usernames := make([]string, len(list.Users))
	for i, user := range list.Users {
		usernames[i] = user.Username
	}

	totalPages := (list.Total + list.Limit - 1) / list.Limit

	response := usernamesListResponse{
		Usernames: usernames,
		Meta: paginationMeta{
			HasPrev:    page > 1,
			HasNext:    page < totalPages,
			Page:       page,
			TotalPages: totalPages,
			Total:      list.Total,
		},
	}

	ctx.JSON(http.StatusOK, response)
}
