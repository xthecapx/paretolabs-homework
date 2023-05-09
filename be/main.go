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

// Second, the endpoint https://searchcaster.xyz/api/profiles?username=dwr
// fetches profile information for a Farcaster user once you know their username (e.g. dwr).
// This endpoint may be helpful to build the profile table.

// - The backend should support two GET requests. One to get a list of **1000** Farcaster users, and one to get basic profile information for a particular Farcaster user. For this exercise, the data can just be kept in memory, though you are welcome to use a database if you prefer.
// - Farcaster has an excellent API that will make your life easier. First, you may use the endpoint [https://api.farcaster.xyz/v2/user?fid=1](https://api.farcaster.xyz/v2/user?fid=1) to fetch a Farcaster user’s username. The `fid` represents a unique identifier (i.e. a Farcaster ID).  ******************************************************************We only want the users with `fid` 1 though 1000.** This endpoint may be helpful to compile a list of Farcaster users.
// - Second, the endpoint [https://searchcaster.xyz/api/profiles?username=dwr](https://searchcaster.xyz/api/profiles?username=dwr) fetches profile information for a Farcaster user once you know their username (e.g. dwr). This endpoint may be helpful to build the profile table.
//     - For a user’s address, please use the `connectedAddress` field from the Farcaster API, not the `address` field.
// - The frontend does not need to be styled exactly as the above screenshot but should at a minimum include two side-by-side sections, with a list of Farcaster users on the left and profile information for a selected user on the right.
// - When a name in the left-side column is clicked, the right-hand information should update accordingly to reflect information about the newly selected user.
// - Try to document your code! Add comments as you go 🙂.

// [{"body":{"id":3,"address":"0x74232BF61e994655592747E20bDF6fa9B9476f79","username":"dwr","displayName":"Dan Romero","bio":"Working on Farcaster and Warpcast.","followers":11093,"following":8550,"avatarUrl":"https://res.cloudinary.com/merkle-manufactory/image/fetch/c_fill,f_png,w_256/https://lh3.googleusercontent.com/MyUBL0xHzMeBu7DXQAqv0bM9y6s4i4qjnhcXz5fxZKS3gwWgtamxxmxzCJX7m2cuYeGalyseCA2Y6OBKDMR06TWg2uwknnhdkDA1AA","isVerifiedAvatar":true,"registeredAt":1620676480000},"connectedAddress":"0xd7029bdea1c17493893aafe29aad69ef892b8ff2"}]
