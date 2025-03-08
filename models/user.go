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
	PasswordHash    string       `json:"passwordHash,omitempty"`
	AvatarURL       string       `json:"avatarUrl"`
	IsActive        bool         `json:"isActive"`
	EmailVerified   bool         `json:"emailVerified"`
	CreatedAt       time.Time    `json:"createdAt"`
	UpdatedAt       time.Time    `json:"updatedAt"`
	OAuthProvider   string       `json:"oauthProvider,omitempty"`
	OAuthPrividerID string       `json:"oauthProviderId,omitempty"`
}
