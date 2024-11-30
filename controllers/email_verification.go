package controllers

import (
	"github.com/gin-gonic/gin"
	"messenger-backend/services"
	"messenger-backend/data-access"
	"net/http"
)

type VerificationRequest struct {
	Email string `json:"email" binding:"required"`
}

func SendVerificationCode(c *gin.Context) {
	email := c.PostForm("email")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	err := services.ProcessVerification(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func VerifyCode(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")

	if email == "" || code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and code are required"})
		return
	}

	isValid, err := dataaccess.VerifyCode(email, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

  if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired verification code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Verification succesfull"})
}