package main
import (
	"github.com/gin-gonic/gin"
	"messenger-backend/routes"
	"messenger-backend/data-access"
	"log"
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
	router.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}