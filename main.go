package main

import (
	"log"
	"productivity-project-backend/controllers"
	"productivity-project-backend/repository"
	"productivity-project-backend/routes"
	"productivity-project-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.InitDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer repository.CloseDatabase(db)

	userRepo := repository.NewUserRepository(db)
	verificationRepo := repository.NewVerificationRepository(db)

	
	authService := services.NewAuthService(userRepo, verificationRepo)

	
	authController := controllers.NewAuthController(authService)
	verificationController := controllers.NewVerificationController(authService)

	router := gin.Default()

	// TODO: Implement auth middleware for logout and other routes

	routes.RegisterRoutes(router, authController, verificationController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
