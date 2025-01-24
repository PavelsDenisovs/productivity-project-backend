package repository

import (
	"fmt"
	"database/sql"
	"errors"
	"messenger-backend/models"
	"time"
)

type UserRepository struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, display_name, email, password_hash, created_at, updated_at, is_active, email_verified)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := r.db.QueryRow(query, user.Username, user.DisplayName, user.Email, user.PasswordHash, time.Now(), time.Now(), true, false).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
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