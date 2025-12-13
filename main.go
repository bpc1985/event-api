package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	db.DB.AutoMigrate(&models.User{}, &models.Event{}, &models.Registration{})

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
