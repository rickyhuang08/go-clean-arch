package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rickyhuang08/gin-project/config"
	"github.com/rickyhuang08/gin-project/delivery/http"
	"github.com/rickyhuang08/gin-project/helpers"
	"github.com/rickyhuang08/gin-project/internal/repository/sql"
	"github.com/rickyhuang08/gin-project/internal/usecase"
	"github.com/rickyhuang08/gin-project/middleware"
	"github.com/rickyhuang08/gin-project/pkg/auth"
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

	// Initialize helper for time abstraction
	timeProvider := helpers.NewRealTimeProvider()

	// Path to public key for JWT validation
	publicKeyPath := cfg.Jwt.PublicKey

	// Initialize middleware module
	mw := middleware.NewMiddlewareModule(timeProvider, publicKeyPath)

	// Apply global middlewares
	mw.RegisterGlobalMiddleware(router)

	// Initialize repository
	userRepo := sql.NewUserRepository()

	// Initialize usecases
	privateKey, err := middleware.LoadPrivateKey(cfg.Jwt.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}

	jwtHelper := auth.NewJWTHelper(timeProvider)
	authUC := usecase.NewAuthUsecase(userRepo, jwtHelper, privateKey)
	userUC := usecase.NewUserUsecase(userRepo)

	// Initialize handler
	handler := http.NewHandler(authUC, userUC)

	// Register routes
	http.RegisterRoutes(router, handler, mw)

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
