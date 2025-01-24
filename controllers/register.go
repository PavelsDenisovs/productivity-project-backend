package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
  "fmt"
  "strings"
  "github.com/go-playground/validator/v10"
	"messenger-backend/models"
	"messenger-backend/services"
	"messenger-backend/utils"
)

type RegisterRequest struct {
	Email       string `json:"email" validate:"required,email"`
	DisplayName string `json:"displayName" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
	}

	// Validate the request fields
	validate := validator.New()
	err := validate.Struct(req)
  if err != nil {
    validationErrors := err.(validator.ValidationErrors)
    var errorMessages []string
    for _, err := range validationErrors {
      field := err.StructField()
      tag := err.Tag()
      errorMessages = append(errorMessages, fmt.Sprintf("Error in field %s: %s validation failed", field, tag))
    }

    c.JSON(http.StatusBadRequest, gin.H{"error": strings.Join(errorMessages, "; ")})
    return
  }

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing user"})
		return
	}

	// Create a new user model
	user := models.User {
		Email: req.Email,
		DisplayName: req.DisplayName,
		Username: req.Username,	
		PasswordHash: hashedPassword,
	}

	// Use the service layer to create the user
	err = services.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "unique constaint") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
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

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}