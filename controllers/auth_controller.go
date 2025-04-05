package controllers

import (
	"net/http"
	"productivity-project-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	GetCurrentUser(c *gin.Context)
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

	if oldSession, err := ac.store.Get(c.Request, "session"); err == nil {
    oldSession.Options.MaxAge = -1
    oldSession.Save(c.Request, c.Writer)
	}

	session, err := ac.store.New(c.Request, "session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Session creation failed"})
		return 
}

	session.Values["user_id"] = user.ID
	session.Values["authenticated"] = true

	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
 
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": user,
	})
}

func (ac *authController) Logout(c *gin.Context) {
	session, err := ac.store.Get(c.Request, "session")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "No active session"})
    return
}
	
	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1
	
	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Logout failed"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func (ac *authController) GetCurrentUser(c *gin.Context) {
	session, err := ac.store.Get(c.Request, "session")
	if err != nil || !session.Values["authenticated"].(bool) {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
    return
  }

	userID := session.Values["user_id"].(uuid.UUID)
  user, err := ac.authService.GetUserByID(userID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"email": user.Email})
}