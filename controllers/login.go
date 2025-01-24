package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"messenger-backend/services"
	"messenger-backend/utils"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
	}

	// Fetch user by email
	user, err := services.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error genereting token"})
		return
	}
	
	// Set the token in a secure httpOnly cookie
	c.SetCookie("refresh_token", token, 7*24*3600, "/", "localhost", false, true)

	// Return the token
	c.JSON(http.StatusOK, gin.H{"token": token})
}