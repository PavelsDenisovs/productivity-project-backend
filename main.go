package main

import (
	"log"
	"messenger-backend/repository"
	"messenger-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	// Initialize PostgreSQL database
	db, err := repository.InitDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer repository.CloseDatabase(db)

	// Set up Gin router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-lenght"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}