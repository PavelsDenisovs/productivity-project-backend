package repository

import (
	"productivity-project-backend/models"
	"time"
)

type NoteRepository interface {
	Create(note *models.Note) error
	GetByDate(userID string, date time.Time)
}