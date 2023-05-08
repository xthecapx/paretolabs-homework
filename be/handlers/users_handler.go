package handlers

import (
	"net/http"
	"strconv"

	"paretolabs-backend/models"

	"github.com/gin-gonic/gin"
)

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
		c.JSON(http.StatusOK, gin.H{"data": []models.User{}, "page": page, "per_page": perPage, "total_users": len(userList)})
		return
	}

	if endIndex > len(userList) {
		endIndex = len(userList)
	}

	// Get the paginated list of users
	paginatedUsers := userList[startIndex:endIndex]

	c.JSON(http.StatusOK, gin.H{"data": paginatedUsers, "page": page, "per_page": perPage, "total_users": len(userList)})
}
