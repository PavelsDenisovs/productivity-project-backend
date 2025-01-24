package repository

import (
	"time"
	"fmt"
	"database/sql"
)

type VerificationRepository struct{
	db *sql.DB
}

func NewVerificationRepository(db *sql.DB) *VerificationRepository {
	return &VerificationRepository{db: db}
}

func (r *VerificationRepository) StoreVerificationCode(email, code string, expiresAt time.Time) error {
	query := `INSERT INTO email_verifications (email, code, expires_at)
	          VALUES ($1, $2, $3) ON CONFLICT (email)
						DO UPDATE SET code = $2, expires_at = $3`
	_, err := db.Exec(query, email, code, expiresAt)
	if err != nil {
		return fmt.Errorf("failed to store verification code: %v", err)
	}
	return nil
}

func (r *VerificationRepository) VerifyCode(email, code string) (bool, error) {
	var exists int
	query := `SELECT 1 FROM email_verifications WHERE email = $1 AND code = $2 AND expires_at > NOW()`
	err := db.QueryRow(query, email, code).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("failed to verify code: %v", err)
	}
	return true, nil
}