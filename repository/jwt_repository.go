package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type JWTTokenRepository interface {
  StoreRefreshToken(userID, refreshToken string, expiresAt time.Time) error
  IsTokenBlacklisted(refreshToken string) (bool, error)
  RevokeToken(refreshToken string) error
}

type jwtTokenRepository struct{
	db *sql.DB
}

func NewJWTTokenRepository(db *sql.DB) JWTTokenRepository {
	return &jwtTokenRepository{db: db}
}

func (r *jwtTokenRepository) StoreRefreshToken(userID string, refreshToken string, expiresAt time.Time) error {
  query := `INSERT INTO jwt_tokens (user_id, refresh_token, is_blacklisted, created_at, expires_at)
            VALUES ($1, $2, $3, $4, $5) ON CONFLICT (refresh_token)
            DO UPDATE SET refresh_token = $2, expires_at = $5, is_blacklisted = $3`
  _, err := r.db.Exec(query, userID, refreshToken, false, time.Now(), expiresAt)
  if err != nil {
    return fmt.Errorf("failed to store refresh token: %v", err)
  }
  return nil
}

func (r *jwtTokenRepository) IsTokenBlacklisted(refreshToken string) (bool, error) {
  var isBlacklisted bool
  query := `SELECT is_blacklisted FROM jwt_tokens WHERE refresh_token = $1`
  err := r.db.QueryRow(query, refreshToken).Scan(&isBlacklisted)
  if err == sql.ErrNoRows {
    return false, nil
  } else if err != nil {
    return false, fmt.Errorf("failed to check token blacklist status: %v", err)
  }
  return isBlacklisted, nil
}

func (r *jwtTokenRepository) RevokeToken(refreshToken string) error {
  query := `UPDATE jwt_tokens SET is_blacklisted = true WHERE refresh_token = $1`
  _, err := r.db.Exec(query, refreshToken)
  if err != nil {
    return fmt.Errorf("failed to revoke token: %v", err)
  }
  return nil
}