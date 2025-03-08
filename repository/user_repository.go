package repository

import (
	"fmt"
	"database/sql"
	"errors"
	"messenger-backend/models"
    "github.com/google/uuid"
	"time"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	// Generate a new UUID for the user ID
	user.ID = uuid.New()

	// Set the default AvatarURL
	user.AvatarURL = "/images/default-profile.svg"

	// Insert the user into the database
	query := `INSERT INTO users (id, username, display_name, email, password_hash, created_at, updated_at, is_active, email_verified)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := r.db.QueryRow(query, user.Username, user.DisplayName, user.Email, user.PasswordHash, time.Now(), time.Now(), true, false).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, display_name, email, password_hash, is_active, email_verified, created_at, updated_at
						FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.DisplayName, &user.Email, &user.PasswordHash, &user.IsActive, &user.EmailVerified, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}
	return &user, nil
}