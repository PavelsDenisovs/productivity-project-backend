package controllers

import (
	"net/http"
	"productivity-project-backend/services"

	"github.com/gin-gonic/gin"
)

type VerificationController interface {
	VerifyEmail(c *gin.Context)
	ResendVerification(c *gin.Context)
}

type verificationController struct {
	authService services.AuthService
}

func NewVerificationController(authService services.AuthService) VerificationController {
	return &verificationController{authService: authService}
}

func (vc *verificationController) VerifyEmail(c *gin.Context) {
	var request struct {
		Email string `json:"email" binding:"required,email"`
		Code  string `json:"code" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err := vc.authService.VerifyEmail(request.Email, request.Code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Verification failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}

func (vc *verificationController) ResendVerification(c *gin.Context) {
	var request struct {
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return 
	}

	err := vc.authService.GenerateAndStoreVerificationCode(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to resend verification"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Verification code resent"})
}