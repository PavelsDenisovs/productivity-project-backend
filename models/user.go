package models

import (
    "time"
    "github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID    `json:"id"`
	Username        string       `json:"username" validate:"required,max=30"`
	DisplayName     string       `json:"displayName" validate:"required,max=50"`
	Email           string       `json:"email" validate:"required,email"`
	PasswordHash    string       `json:"-"`
	AvatarURL       string       `json:"avatarUrl"`
	IsActive        bool         `json:"isActive"`
	EmailVerified   bool         `json:"emailVerified"`
	CreatedAt       time.Time    `json:"createdAt"`
	UpdatedAt       time.Time    `json:"updatedAt"`
	OAuthProvider   string       `json:"oauthProvider,omitempty"`
	OAuthProviderID string       `json:"oauthProviderId,omitempty"`
}

type UserUpdate struct {
    Username    string `json:"username,omitempty" validate:"omitempty,max=30"`
    DisplayName string `json:"displayName,omitempty" validate:"omitempty,max=50"`
    AvatarURL   string `json:"avatarUrl,omitempty" validate:"omitempty,url"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required,min=8"`
}