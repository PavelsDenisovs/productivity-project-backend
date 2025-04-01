package main

import (
	"log"
	"net/http"
	"productivity-project-backend/controllers"
	"productivity-project-backend/repository"
	"productivity-project-backend/routes"
	"productivity-project-backend/services"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var store *sessions.CookieStore

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

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

	
	authService := services.NewAuthService(userRepo, verificationRepo)

	
	authController := controllers.NewAuthController(authService, store)
	verificationController := controllers.NewVerificationController(authService, store)

	router := gin.Default()

	// TODO: Implement auth middleware for logout and other routes

	routes.RegisterRoutes(router, authController, verificationController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
