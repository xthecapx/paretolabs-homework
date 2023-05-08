package main

import (
	"paretolabs-backend/handlers"
	"paretolabs-backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	userListMap, userList := services.GetUsersList()
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		handlers.GetUsersHandler(c, userList)
	})
	router.GET("/users/:fid", func(c *gin.Context) {
		handlers.GetUserDetailsHandler(c, userListMap)
	})
	router.Run(":1234")
}
