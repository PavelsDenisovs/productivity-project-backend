package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"productivity-project-backend/controllers"
	"productivity-project-backend/repository"
	"productivity-project-backend/routes"
	"productivity-project-backend/services"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var store *sessions.CookieStore

func main() {
	// TODO: implement .env.production and .env.development separation
	if os.Getenv("RENDER") == "" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, continuing with system env...")
		}
	}

	
  gob.Register(uuid.UUID{})

	store  = sessions.NewCookieStore(
		[]byte(os.Getenv("SESSION_SECRET")),
	)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   30 * 86400,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	db, err := repository.InitDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer repository.CloseDatabase(db)

	userRepo := repository.NewUserRepository(db)
	verificationRepo := repository.NewVerificationRepository(db)
	noteRepo := repository.NewNoteRepository(db)

	
	authService := services.NewAuthService(userRepo, verificationRepo)
	noteService := services.NewNoteService(noteRepo)

	
	authController := controllers.NewAuthController(authService, store)
	verificationController := controllers.NewVerificationController(authService, store)
	noteController := controllers.NewNoteController(noteService)

	router := gin.Default()

	// TODO: Implement auth middleware for logout and other routes

	routes.RegisterRoutes(router, authController, verificationController, noteController, store)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
