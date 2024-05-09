package main

import (
	"strconv"

	"example.com/db"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()
	server := gin.Default()

	server.GET("/events",getEvents)
	server.GET("/events/:eventId",getEvent)
	server.POST("/events",createEvent)

	server.Run(":8080") //localhost:8080

}

func getEvents(context *gin.Context){
	events, err := models.GetAllEvents()
	if(err != nil){
		context.JSON(500, gin.H{"message: ": "could not get events"})
		return
	}
	context.JSON(200, events)
}

func getEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("eventId"),10,64)

	if err != nil {
		context.JSON(400, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(500, gin.H{"message": "Could not get event"})
		return
	}

	context.JSON(200,event)


}

func createEvent(context *gin.Context) {
	event := models.Event{}
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserId = 1

	err = event.Save()
	if err != nil {
		context.JSON(500, gin.H{"message: ": "could not create event"})
		return
	}

	context.JSON(201, gin.H{"msg":"Event created!","data":gin.H{"event": event}})
}