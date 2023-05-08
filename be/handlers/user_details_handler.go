package handlers

import (
	"net/http"
	"strconv"

	"paretolabs-backend/models"

	"github.com/gin-gonic/gin"
)

func GetUserDetailsHandler(c *gin.Context, userList []models.User) {
	// Get the FID parameter from the request URL
	fidStr := c.Param("fid")

	// Convert the FID to an integer
	fid, err := strconv.Atoi(fidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid FID"})
		return
	}

	// Find the user with the matching FID in the userList
	var userDetails *models.User
	for _, user := range userList {
		if user.FID == fid {
			userDetails = &user
			break
		}
	}

	if userDetails == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": userDetails})
	}
}
