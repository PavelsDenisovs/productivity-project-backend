package services

import (
	"productivity-project-backend/models"
	"productivity-project-backend/repository"

	"github.com/google/uuid"
)

type NoteService interface {
	GetAllNotes(userID uuid.UUID) ([]models.Note, error)
	CreateNote(note *models.Note) error
	UpdateNote(note *models.Note) error
}

type noteService struct {
	noteRepo repository.NoteRepository
}

func NewNoteService(noteRepo repository.NoteRepository) NoteService {
	return &noteService{noteRepo: noteRepo}
}

func (s *noteService) GetAllNotes(userID uuid.UUID) ([]models.Note, error) {
	return s.noteRepo.GetAllNotes(userID)
}

func (s *noteService) CreateNote(note *models.Note) error {
	return s.noteRepo.Create(note)
}

func (s *noteService) UpdateNote(note *models.Note) error {
	return s.noteRepo.Update(note)
}