package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"messenger-backend/utils"
	"messenger-backend/services"
)

type ValidationRequest struct {
	DisplayName string `json:"displayName,omitempty"`
  Email       string `json:"email,omitempty"`
  Username    string `json:"username,omitempty"`
  Password    string `json:"password,omitempty"`
}

func Validate(c *gin.Context) {
	var req ValidationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	errors := make(map[string]string)

  if req.DisplayName != "" {
    if msg := utils.ValidateDisplayName(req.DisplayName); msg != "" {
      errors["displayName"] = msg
    }
  }

  if req.Email != "" {
    if msg := utils.ValidateEmail(req.Email); msg != "" {
      errors["email"] = msg
    } else {
      exists, err := services.IsEmailInUse(req.Email)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
				return
			}
			if exists {
				errors["email"] = "Email is already in use"
			}
    }
  }

  if req.Username != "" {
    if msg := utils.ValidateUsername(req.Username); msg != "" {
      errors["username"] = msg
    } else {
      exists, err := services.IsUsernameInUse(req.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
				return
			}
			if exists {
				errors["username"] = "Username is already in use"
			}
    }
  }

  if req.Password != "" {
    if msg := utils.ValidatePassword(req.Password); msg != "" {
      errors["password"] = msg
    }
  }

	c.JSON(http.StatusOK, gin.H{"errors": errors})
}