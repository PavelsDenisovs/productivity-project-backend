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
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
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

	// Ensure passwords match
	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing user"})
		return
	}

	// Create a new user model
	user := models.User {
		Username: req.Username,
		Email: req.Email,
		PasswordHash: hashedPassword,
	}

	// Use the service layer to create the user
	err = services.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}