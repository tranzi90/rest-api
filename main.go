package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"msg": "error"})
		return
	}

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"msg": "Event Created", "event": event})
}
