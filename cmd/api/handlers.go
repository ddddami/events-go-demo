package main

import (
	"net/http"

	"github.com/ddddami/events-go-demo/internal/models"
	"github.com/gin-gonic/gin"
)

func healthcheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"mode":    gin.Mode(),
		"version": version,
	})
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	events := models.GetAllEvents()
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data\n", "error": err})
		return
	}

	event.ID = len(events) + 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}
