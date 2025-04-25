package main

import "github.com/gin-gonic/gin"

func registerRoutes(server *gin.Engine) {
	server.GET("healthcheck/", healthcheck)

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
}
