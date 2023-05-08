package handlers

import (
	"net/http"
	"strconv"

	"paretolabs-backend/models"

	"github.com/gin-gonic/gin"
)

func GetUserDetailsHandler(c *gin.Context, userList map[int]models.User) {
	// Get the FID parameter from the request URL
	fidStr := c.Param("fid")

	// Convert the FID to an integer
	fid, err := strconv.Atoi(fidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid FID"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid FID"})
		return
	}

	user, ok := userList[fid]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
