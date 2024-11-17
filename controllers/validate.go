package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"messenger-backend/utils"
	"messenger-backend/services"
)

type ValidationRequest struct {
	Fields []utils.FieldValidationRequest `json:"fields"`
}

func Validate(c *gin.Context) {
	var req ValidationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	errors := utils.ValidateFields(req.Fields)

	for _, field := range req.Fields {
		if field.FieldName == "email" && errors[field.FieldName] == "" {
			exists, err := services.IsEmailInUse(field.Value)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
				return
			}
			if exists {
				errors[field.FieldName] = "Email is already in use"
			}
		}
		if field.FieldName == "username" && errors[field.FieldName] == "" {
			exists, err := services.IsUsernameInUse(field.Value)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
				return
			}
			if exists {
				errors[field.FieldName] = "Username is already in user"
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"errors": errors})
}