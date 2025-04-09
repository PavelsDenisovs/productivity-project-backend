package repository

import (
	"database/sql"
	"productivity-project-backend/models"
	"time"

	"github.com/google/uuid"
)

type NoteRepository interface {
	GetAllNotes(userID uuid.UUID) ([]models.Note, error)
	Create(note *models.Note) error
	Update(note *models.UpdateNoteDTO) error
}

type noteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) NoteRepository {
	return &noteRepository{db: db}
}

func (r *noteRepository) GetAllNotes(userID uuid.UUID) ([]models.Note, error) {
	rows, err := r.db.Query(`
		SELECT id, user_id, date, content, sleep_quality, created_at, updated_at
		FROM notes WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var notes []models.Note
	for rows.Next() {
		var note models.Note
		err := rows.Scan(
			&note.ID,
			&note.UserID,
			&note.Date,
			&note.Content,
			&note.SleepQuality,
			&note.CreatedAt,
			&note.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
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

func (r *noteRepository) Update(noteData *models.UpdateNoteDTO) error {
	_, err := r.db.Exec(`
		UPDATE notes SET content = $1, sleep_quality = $2, updated_at = $3
		WHERE id = $4`, noteData.Content, noteData.SleepQuality, time.Now(), noteData.ID)
	return err
}