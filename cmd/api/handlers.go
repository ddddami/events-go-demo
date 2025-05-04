package main

import (
	"errors"
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

	app.events.Insert(e.Title, e.Description, e.Location, e.DateTime, e.UserID)
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": e})
}

func (app *application) updateEvent(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID", "error": err})
		return
	}

	_, err = app.events.GetByID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "No event with the given ID", "error": err})
		return
	}

	var newEvent models.Event

	err = context.ShouldBind(&newEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	newEvent.ID = id
	err = app.events.Update(&newEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not update the event", "err": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated"})
}

func (app *application) deleteEvent(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find an event with the given ID"})
		return
	}
	_, err = app.events.GetByID(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not find an event with the given id"})
		return
	}
	err = app.events.Delete(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not delete the event"})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{"message": "Event deleted"})
}

func (app *application) saveUser(context *gin.Context) {
	var u models.User
	err := context.ShouldBindJSON(&u)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	err = app.users.Register(u.Email, u.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not create user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created", "user": u})
}

func (app *application) login(context *gin.Context) {
	var u models.User
	err := context.ShouldBindJSON(&u)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	id, err := app.users.Authenticate(u.Email, u.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Email or password is incorrect"})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Trouble authenticating user", "err": err.Error()})
		}
		return
	}

	token, err := app.users.GenerateToken(u.Email, id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Trouble authenticating user", "err": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
