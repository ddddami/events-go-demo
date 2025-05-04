package main

import "github.com/gin-gonic/gin"

func registerRoutes(server *gin.Engine, app *application) {
	authenticated := server.Group("/")
	authenticated.Use(app.authenticate)

	server.GET("healthcheck/", healthcheck)

	server.GET("/events", app.getEvents)
	server.GET("/events/:id", app.getEvent)
	authenticated.PUT("/events/:id", app.updateEvent)
	authenticated.DELETE("/events/:id", app.deleteEvent)
	authenticated.POST("/events", app.authenticate, app.createEvent)
	// users
	server.POST("/signup", app.saveUser)
	server.POST("/login", app.login)
}
