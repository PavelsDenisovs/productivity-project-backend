package controllers

import (
	"net/http"
	"productivity-project-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type authController struct {
	authService services.AuthService
	store       *sessions.CookieStore
}

func NewAuthController(authService services.AuthService, store *sessions.CookieStore) AuthController {
	return &authController{
		authService: authService,
		store: store,
	}
}

func (ac *authController) Register(c *gin.Context) {
	var request struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	user, err := ac.authService.Register(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Registration failed"})
		return
	}

	if err := ac.authService.GenerateAndStoreVerificationCode(request.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send verification email"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful. Check your email for verification",
		"user_id": user.ID,
	})
}

func (ac *authController) Login(c *gin.Context) {
	var request struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	user, err := ac.authService.Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": user,
	})
}