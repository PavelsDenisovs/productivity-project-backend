package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID           uuid.UUID `json:"id" db:"id"`
  UserID       uuid.UUID `json:"user_id" db:"user_id"`
	Date         time.Time `json:"-" db:"date"`
	Content      string    `json:"content" db:"content"`
	SleepQuality *int      `json:"sleep_quality,omitempty" db:"sleep_quality"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

func (n *Note) MarshalJSON() ([]byte, error) {
	type Alias Note
	return json.Marshal(&struct {
		Date string `json:"date"`
		*Alias
	}{
		Date:  n.Date.Format("2006-01-02"),
		Alias: (*Alias)(n),
	})
}