package handlers

import (
	"net/http"
	"strconv"

	"paretolabs-backend/models"

	"github.com/gin-gonic/gin"
)

// GetUsersHandler returns a paginated list of users.
//
//	@Summary		Get paginated list of users
//	@Description	Get a paginated list of users with pagination support.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			page		query		int	false	"Page number"				default(1)
//	@Param			per_page	query		int	false	"Number of users per page"	default(15)
//	@Success		200			{object}	models.PaginatedUsersResponse
//	@Failure		400			{object}	models.ErrorResponse
//	@Router			/users [get]
func GetUsersHandler(c *gin.Context, userList []models.User) {
	// Get the pagination parameters from the query string
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "15")

	// Convert the pagination parameters to integers
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid per_page parameter"})
		return
	}

	// Calculate the start and end indexes for the pagination
	startIndex := (page - 1) * perPage
	endIndex := startIndex + perPage

	if startIndex >= len(userList) {
		c.JSON(http.StatusOK, gin.H{"userList": []models.User{}, "page": page, "per_page": perPage, "total_count": len(userList)})
		return
	}

	if endIndex > len(userList) {
		endIndex = len(userList)
	}

	// Get the paginated list of users
	paginatedUsers := userList[startIndex:endIndex]

	c.JSON(http.StatusOK, gin.H{"userList": paginatedUsers, "page": page, "per_page": perPage, "total_count": len(userList)})
}
