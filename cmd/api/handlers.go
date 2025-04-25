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

