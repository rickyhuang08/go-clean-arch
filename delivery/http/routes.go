package http

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rickyhuang08/gin-project/middleware"
)

// RegisterRoutes sets up API routes
func RegisterRoutes(r *gin.Engine, handler *Handler, mw *middleware.MiddlewareModule) {
	r.POST("/login", handler.LoginHandler)

	auth := r.Group("/api/v1")
	if err := mw.RegisterAuthMiddleware(auth); err != nil {
		log.Fatalf("Failed to register auth middleware: %v", err)
	}
	auth.GET("/profile", handler.ProfileHandler)
}