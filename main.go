package main

import (
	"log"
	"messenger-backend/data-access"
	"messenger-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize PostgreSQL database
	err := dataaccess.InitializeDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Set up Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}