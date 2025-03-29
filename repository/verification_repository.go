package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)


type VerificationRepository interface {
	StoreVerificationCode(userID uuid.UUID, code string) error
	VerifyCode(userID uuid.UUID, code string) (bool, error)
}

type verificationRepository struct{
	db *sql.DB
}

func NewVerificationRepository(db *sql.DB) VerificationRepository {
	return &verificationRepository{db: db}
}

func (r *verificationRepository) StoreVerificationCode(userID uuid.UUID, code string) error {
	expiresAt := time.Now().Add(time.Minute * 15)
	query := `INSERT INTO verification_codes (user_id, code, expires_at)
	          VALUES ($1, $2, $3) 
						ON CONFLICT (user_id)
						WHERE NOT used AND expires_at > NOW()
						DO UPDATE SET 
							code = EXCLUDED.code, 
							expires_at = EXCLUDED.expires_at,
							used = false`
	_, err := r.db.Exec(query, userID, code, expiresAt)
	if err != nil {
		return fmt.Errorf("failed to store verification code: %v", err)
	}
	return nil
}

func (r *verificationRepository) VerifyCode(userID uuid.UUID, code string) (bool, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	var valid bool
	err = tx.QueryRow(`
		UPDATE verification_codes
		SET used = true
		WHERE user_id = $1
			AND code = $2
			AND expires_at > NOW()
			AND NOT used
		RETURNING true`,
		userID, code,
	).Scan(&valid)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return valid, tx.Commit()
}