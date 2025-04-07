package models

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID           uuid.UUID    `json:"id" db:"id"`
	Date         string    `json:"date" db:"date"`
	Content      string    `json:"content" db:"content"`
	SleepQuality *int      `json:"sleep_quality,omitempty" db:"sleep_quality"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}