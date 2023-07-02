package main

import (
	"fmt"

	"github.com/athunlal/api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load configuration: %s\n", err.Error())
		return
	}

	// Initialize database
	db, err := database.InitDB(cfg.DBUrl)
	if err != nil {
		fmt.Printf("Failed to initialize database: %s\n", err.Error())
		return
	}
	defer db.Close()

	router := gin.Default()

	// Routes setup
	routes.Setup(router)

	// Run the server
	router.Run(cfg.Port)
}
