package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rickyhuang08/gin-project/internal/entity"
)

// LoginHandler processes user login
func (h *Handler) LoginHandler(c *gin.Context) {
	var req entity.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	token, err := h.AuthUsecase.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}