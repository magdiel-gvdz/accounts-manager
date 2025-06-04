package main

import (
	"log"

	"github.com/Magdiel-GVdz/accounts-manager/database"
	"github.com/Magdiel-GVdz/accounts-manager/internal/adapters/repository"
	"github.com/Magdiel-GVdz/accounts-manager/internal/application"
	"github.com/Magdiel-GVdz/accounts-manager/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	repo := repository.NewGormUserRepository(database.DB)
	service := application.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	router := gin.Default()
	handler.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
