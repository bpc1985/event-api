package routes

import (
	"time"

	"example.com/rest-api/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// server.GET("/events", getEvents)    // GET, POST, PUT, PATCH, DELETE
	// server.GET("/events/:id", getEvent) // /events/1, /events/5

	// authenticated := server.Group("/")
	// authenticated.Use(middlewares.Authenticate)
	// authenticated.POST("/events",createEvent)
	// authenticated.PUT("/events/:id", updateEvent)
	// authenticated.DELETE("/events/:id", deleteEvent)
	// authenticated.POST("/events/:id/register", registerForEvent)
	// authenticated.DELETE("/events/:id/register", cancelRegistration)

	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	server.Use(cors.New(config))

	index := server.Group("/")
	{
		index.GET("/events", getEvents)
		index.GET("/events/:id", getEvent)

		index.GET("/events/:id/attendees", getAttendeesForEvent)
		index.GET("/attendees/:id/events", getEventsByAttendee)

		index.POST("/signup", signup)
		index.POST("/login", login)
	}

	authGroup := index.Group("/events")
	authGroup.Use(middlewares.Authenticate)
	{
		authGroup.POST("", createEvent)
		authGroup.PUT("/:id", updateEvent)
		authGroup.DELETE("/:id", deleteEvent)
		authGroup.POST("/:id/register", registerForEvent)
		authGroup.DELETE("/:id/register", cancelRegistration)
	}
}
