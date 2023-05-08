package main

import (
	"log"
	"os"

	"paretolabs-backend/handlers"
	"paretolabs-backend/models"
	"paretolabs-backend/services"
	"paretolabs-backend/utils"

	"github.com/gin-gonic/gin"
)

func getUserList() []models.User {
	if _, err := os.Stat(utils.GetUserFilePath()); os.IsNotExist(err) {
		// users.json does not exist, fetch data and write to file
		userList, err := services.GetDataAndWriteToFile(1, 10)
		if err != nil {
			log.Fatalf("Failed to fetch data and write to file: %v", err)
		}

		return userList
	}
	// users.json exists, read data from file
	userList, err := services.ReadUsersFromFile()
	if err != nil {
		log.Fatalf("Failed to read users from file: %v", err)
	}

	return userList
}

func main() {
	userList := getUserList()
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		handlers.GetUsersHandler(c, userList)
	})
	router.GET("/users/:fid", func(c *gin.Context) {
		handlers.GetUserDetailsHandler(c, userList)
	})
	router.Run(":1234")
}
