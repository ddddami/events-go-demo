package main

import (
	"net/http"
	"strconv"

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

func (app *application) getEvents(context *gin.Context) {
	events, err := app.events.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func (app *application) getEvent(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}

	event, err := app.events.GetByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func (app *application) createEvent(context *gin.Context) {
	var e models.Event

	err := context.ShouldBindJSON(&e)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data\n", "error": err})
		return
	}

	e.UserID = 1
	app.events.Insert(e.Title, e.Description, e.Location, e.DateTime, e.UserID)
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": e})
}
