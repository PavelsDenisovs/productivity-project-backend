package controllers

import (
	"net/http"
	"productivity-project-backend/models"
	"productivity-project-backend/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NoteController interface {
	GetAllNotes(c *gin.Context)
	CreateNote(c *gin.Context)
	UpdateNote(c *gin.Context)
}

type noteController struct {
	noteService services.NoteService
}

func NewNoteController(noteService services.NoteService) NoteController {
	return &noteController{noteService: noteService}
}

func (nc *noteController) GetAllNotes(c *gin.Context) {
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in session"})
		return
	}
	userID := userIDVal.(uuid.UUID)

	notes, err := nc.noteService.GetAllNotes(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notes": notes})
}

func (nc *noteController) CreateNote(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	// User ID from session
	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in session"})
		return
	}
	note.UserID = userIDVal.(uuid.UUID)

	if note.Content == "" {
		note.Content = ""
	}

	note.Date = time.Date(
		note.Date.Year(), 
		note.Date.Month(), 
		note.Date.Day(), 
		0, 0, 0, 0, 
		time.UTC,
	)

	if err := nc.noteService.CreateNote(&note); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"note": note})
}

func (nc *noteController) UpdateNote(c *gin.Context) {
	var noteData models.UpdateNoteDTO
	if err := c.ShouldBindJSON(&noteData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	noteID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}
	noteData.ID = noteID

	if err := nc.noteService.UpdateNote(&noteData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"note": noteData})
}