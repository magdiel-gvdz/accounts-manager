package main

import (
	"log"

	"github.com/Magdiel-GVdz/accounts-manager/database"
	"github.com/Magdiel-GVdz/accounts-manager/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	database.Connect()

	// Create a new Gin router
	router := gin.Default()

	// Set up routes
	routes.UserRoutes(router)

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
