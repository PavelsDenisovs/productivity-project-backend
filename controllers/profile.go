package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"messenger-backend/services"
)

func Profile(c *gin.Context) {
	// Assume userID is available from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Fetch user by ID
	user, err := services.GetUserById(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Respond with user's profile
	profile := map[string]interface{}{
		"id": user.ID,
		"username": user.Username,
		"email": user.Email,
		"avatarUrl": user.AvatarURL,
	}
	c.JSON(http.StatusOK, profile)
}