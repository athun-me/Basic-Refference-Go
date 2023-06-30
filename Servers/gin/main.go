package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define routes and their handlers
	router.GET("/", helloWorld)
	router.GET("/user", getUser)

	// Start the server
	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}

func helloWorld(c *gin.Context) {
	c.String(200, "Hello, World!")
}

func getUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": "This is JSON response",
	})
}
