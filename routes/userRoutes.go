package routes

import (
	"github.com/Magdiel-GVdz/accounts-manager/controllers"
	"github.com/Magdiel-GVdz/accounts-manager/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	// Rutas públicas
	router.POST("/users", controllers.CreateUser)
	router.POST("/login", controllers.LoginUser)

	// Rutas protegidas
	auth := router.Group("/users")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/:id", controllers.GetUser)
		auth.PUT("/:id", controllers.UpdateUser)
		auth.DELETE("/:id", controllers.DeleteUser)
		auth.GET("/", controllers.ListUsers)
	}
}
