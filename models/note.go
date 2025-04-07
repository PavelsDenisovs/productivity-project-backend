package models

import "time"

type Note struct {
	ID           string    `json:"id" db:"id"`
	Date         string    `json:"date" db:"date"`
	Content      string    `json:"content" db:"content"`
	SleepQuality *int      `json:"sleep_quality,omitempty" db:"sleep_quality"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}