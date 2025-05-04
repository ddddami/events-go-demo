package main

import "github.com/gin-gonic/gin"

func registerRoutes(server *gin.Engine, app *application) {
	server.GET("healthcheck/", healthcheck)

	server.GET("/events", app.getEvents)
	server.GET("/events/:id", app.getEvent)
	server.POST("/events", app.createEvent)
}
