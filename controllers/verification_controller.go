package controllers

import (
	"net/http"
	"productivity-project-backend/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type VerificationController interface {
	VerifyEmail(c *gin.Context)
	ResendVerification(c *gin.Context)
}

type verificationController struct {
	authService services.AuthService
	store       *sessions.CookieStore
}

func NewVerificationController(authService services.AuthService, store *sessions.CookieStore) VerificationController {
	return &verificationController{
		authService: authService,
		store:       store,
	}
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

	user, err := vc.authService.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User retrieval failed"})
		return
	}

  if oldSession, err := vc.store.Get(c.Request, "session"); err == nil {
    oldSession.Options.MaxAge = -1
    oldSession.Save(c.Request, c.Writer)
	}

	session, err := vc.store.New(c.Request, "session")
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