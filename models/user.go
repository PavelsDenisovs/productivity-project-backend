package models

import "time"

type User struct {
	ID              uint      `json:"id"`
	Username        string    `json:"username" validate:"required"`
	AvatarURL       string    `json:"avatarUrl"`
	Email           string    `json:"email" validate:"required,email"`
	OAuthProvider   string    `json:"oauthProvider,omitempty"`
	OAuthPrividerID string    `json:"oauthProviderId,omitempty"`
	PasswordHash    string    `json:"passwordHash,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
