package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"paretolabs-backend/models"

	"github.com/gin-gonic/gin"
)

// GetUserDetailsHandler returns details of a specific user.
// @Summary Get user details
// @Description Get details of a user based on the FID.
// @Tags Users
// @Accept json
// @Produce json
// @Param fid path int true "User FID"
// @Success 200 {object} models.UserProfileResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users/{fid} [get]
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

	url := "https://searchcaster.xyz/api/profiles?username=" + user.Username

	log.Print(url)
	// Send a GET request to the URL to fetch the data
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to fetch data for fid %d: %v", fid, err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Failed to read response body for fid %d %s: %v", fid, user.Username, err)
	}

	// Parse the received data into the struct
	var profile []models.ProfileResponse
	err = json.Unmarshal(body, &profile)
	if err != nil {
		log.Printf("Failed to parse data into struct for fid %d: %v", fid, err)
	}

	log.Print(profile[0].Body)
	userInformation := models.UserProfile{
		UserInformation:    user,
		Address:            profile[0].ConnectedAddress,
		ProfileInformation: profile[0].Body,
	}

	// c.JSON(http.StatusOK, userInformation)
	c.JSON(http.StatusOK, gin.H{"profileData": userInformation})
}
