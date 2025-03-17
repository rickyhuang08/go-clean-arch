package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProfileHandler returns user profile (mocked)
func (h *Handler) ProfileHandler(c *gin.Context) {
	userID := 1 // Simulating authenticated user
	user, err := h.UserUsecase.GetUserProfile(userID)
	if err != nil || user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}