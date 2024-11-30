package dataaccess

import (
	"messenger-backend/models"
  "os"
  "github.com/joho/godotenv"
	"database/sql"
	"errors"
	"log"
	"time"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	envFile := ".env." + env
  if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Error loading %s file: %v", envFile, err)
	}
}

// InitializeDB: Initializes the PostgreSQL connection
func InitializeDB() error {
	Init()

  connectionString := os.Getenv("DB_URL")
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	return db.Ping()
}

// SaveUser: Inserts a new user into the database
func SaveUser(user models.User) error {
	query := `INSERT INTO users (username, email, password_hash, avatar_url, created_at, updated_at)
						VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.QueryRow(query, user.Username, user.Email, user.PasswordHash, user.AvatarURL, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

// FindUserByEmail: Finds a user by their email
func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password_hash, avatar_url, created_at, updated_at FROM users WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.AvatarURL, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

// FindUserByUsername: Finds a user by their username
func FindUserByUsername(username string) (models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password_hash, avatar_url, created_at, updated_at FROM users WHERE email = $1`
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.AvatarURL, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

// FindUserByID: Finds a user by their ID
func FindUserByID(userID uint) (models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password_hash, avatar_url, created_at, updated_at FROM users WHERE id = $1`
	err := db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.AvatarURL, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func StoreVerificationCode(email, code string, expiration time.Time) error {
  query := `
    INSERT INTO email_verifications (email, code, expires_at)
    VALUES ($1, $2, $3)
    ON CONFLICT (email)
    DO UPDATE SET code = $2, expires_at = $3
  `
	
  _, err := db.Exec(query, email, code, expiration)
  if err != nil {
      return err
  }
  return nil
}

func VerifyCode(email, code string) (bool, error) {
	query := `
		SELECT 1 FROM email_verifications
		WHERE email = $1 AND code = $2 AND expires_at > NOW()
	`
	var exists int
	err := db.QueryRow(query, email, code).Scan(&exists)
    if err == sql.ErrNoRows {
        return false, nil
    } else if err != nil {
        return false, err
    }
    return true, nil 
}