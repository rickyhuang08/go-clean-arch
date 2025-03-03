package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rickyhuang08/gin-project/config"
)

func startServer() error {
	// Load configuration
	cfg, err := config.NewConfig()
	log.Printf("config : %+v", cfg)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Start server
	log.Printf("Server running on port %s", cfg.Server.Port)
	return router.Run(":" + cfg.Server.Port)
}

func main() {
	// Handle errors properly
	if err := startServer(); err != nil {
		log.Fatal(err) // Only main.go decides to stop the app
	}
}
