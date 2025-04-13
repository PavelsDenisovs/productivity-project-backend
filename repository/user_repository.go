package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"productivity-project-backend/models"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	MarkEmailAsVerified(email string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	// Generate a new UUID for the user ID
	user.ID = uuid.New()

	// Insert the user into the database
	query := `INSERT INTO users (id, email, password_hash, created_at, updated_at, is_verified)
						VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := r.db.QueryRow(query, user.ID, user.Email, user.PasswordHash, time.Now(), time.Now(), false).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password_hash, is_verified, created_at, updated_at
						FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	query := `SELECT id, email, password_hash, is_verified, created_at, updated_at
						FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found: %v", id)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user by id: %v", err)
	}
	return &user, nil
}

func (r *userRepository) MarkEmailAsVerified(email string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`UPDATE users SET is_verified = true WHERE email = $1`, email)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM verification_codes WHERE user_id=(
		SELECT id FROM users WHERE email = $1
	)`, email)
	if err != nil {
		return err
	}
	return tx.Commit()
}

