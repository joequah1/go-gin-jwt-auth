package main

import (
	"github.com/gin-gonic/gin"
	"ginjwtauth/models"
	"ginjwtauth/controllers"
	"ginjwtauth/middlewares"
)

func main() {
	// Initialize the Gin router
	r := gin.Default()

	// Load database connection
	models.ConnectDataBase()

	// Group routes
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			// Register and login routes
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// Profile route with JWT authentication middleware
		api.GET("/profile", middlewares.JwtAuthMiddleware(), controllers.Profile)
	}

	// Run the server
	r.Run(":8080")
}
