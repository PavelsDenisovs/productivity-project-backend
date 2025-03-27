package main

import (
	"log"
	"productivity-project-backend/repository"
	"productivity-project-backend/routes"

	"github.com/gin-gonic/gin"
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

	// Temporary solution
	authMiddleware := func(c *gin.Context) {
    // Empty middleware for now
    c.Next()
}
	// Register routes
	routes.RegisterRoutes(router, authMiddleware)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
