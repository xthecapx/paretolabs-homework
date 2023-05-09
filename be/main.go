package main

import (
	"paretolabs-backend/handlers"
	"paretolabs-backend/services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "paretolabs-backend/docs"

	"github.com/gin-gonic/gin"
)

//	@title			Pareto labs - Homework
//	@version		1.0
//	@description	Solution to the test

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	ing.cristian.marquez@gmail.com

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

//	@host						localhost:1234
//	@BasePath					/
//	@query.collection.format	multi
func main() {
	userListMap, userList := services.GetUsersList()
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		handlers.GetUsersHandler(c, userList)
	})
	router.GET("/users/:fid", func(c *gin.Context) {
		handlers.GetUserDetailsHandler(c, userListMap)
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":1234")
}
