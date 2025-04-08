package repository

import (
	"database/sql"
	"productivity-project-backend/models"
	"time"

	"github.com/google/uuid"
)

type NoteRepository interface {
	Create(note *models.Note) error
	GetByDate(userID uuid.UUID, date time.Time) (*models.Note, error)
	Update(note *models.Note) error
}

type noteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) NoteRepository {
	return &noteRepository{db: db}
}

func (r *noteRepository) Create(note *models.Note) error {
	note.ID = uuid.New()
	now := time.Now()
	note.CreatedAt = now
	note.UpdatedAt = now
	_, err := r.db.Exec(`
		INSERT INTO notes (id, user_id, date, content, sleep_quality, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, note.ID, note.UserID, note.Date, note.Content, note.SleepQuality, note.CreatedAt, note.UpdatedAt)
	return err
}

func (r *noteRepository) GetByDate(userID uuid.UUID, date time.Time) (*models.Note, error) {
	var note models.Note
	err := r.db.QueryRow(`
		SELECT id, user_id, date, content, sleep_quality, created_at, updated_at
		FROM notes WHERE user_id = $1 AND date = $2
	`, userID, date).Scan(
		&note.ID, 
		&note.UserID,
		&note.Content,
		&note.SleepQuality,
		&note.CreatedAt,
	  &note.UpdatedAt,
	)
	return &note, err
}

func (r *noteRepository) Update(note *models.Note) error {
	note.UpdatedAt = time.Now()
	_, err := r.db.Exec(`
		UPDATE notes SET content = $1, sleep_quality = $2, updated_at = $3
		WHERE id = $4`, note.Content, note.SleepQuality, note.UpdatedAt, note.ID)
	return err
}