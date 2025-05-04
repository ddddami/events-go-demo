package main

import "github.com/gin-gonic/gin"

func registerRoutes(server *gin.Engine, app *application) {
	server.GET("healthcheck/", healthcheck)

	server.GET("/events", app.getEvents)
	server.GET("/events/:id", app.getEvent)
	server.PUT("/events/:id", app.updateEvent)
	server.DELETE("/events/:id", app.deleteEvent)
	server.POST("/events", app.createEvent)
	// users
	server.POST("/signup", app.saveUser)
}
