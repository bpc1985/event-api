package main

import (
	_ "example.com/rest-api/docs"
	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

// @title Event Management API
// @version 1.0
// @description REST API for managing events with user authentication
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db.InitDB()
	db.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Registration{})

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
