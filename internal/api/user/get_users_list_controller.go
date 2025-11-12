package user_api

import (
	"net/http"
	"strconv"

	user_application "main/internal/application/user/usecases"

	"github.com/gin-gonic/gin"
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
	Data []usernameData `json:"data"`
	Meta paginationMeta `json:"meta"`
}

type usernameData struct {
	Username string `json:"username"`
	ID       string `json:"id"`
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

	usernamesData := make([]usernameData, len(list.Users))
	for i, user := range list.Users {
		usernamesData[i] = usernameData{
			Username: user.Username,
			ID:       user.ID,
		}
	}

	totalPages := (list.Total + list.Limit - 1) / list.Limit

	response := usernamesListResponse{
		Data: usernamesData,
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
