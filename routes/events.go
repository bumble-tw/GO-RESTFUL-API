package routes

import (
	"strconv"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

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

func updateEvent(context *gin.Context) {
    // 從 URL 參數中解析事件 ID，例如從 "/events/:eventId"。
    // strconv.ParseInt 嘗試將字符串轉換為 int64，基數設為 10。
    eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
    if err != nil {
        context.JSON(400, gin.H{"message": "Could not parse event id"})
        return
    }

    // 試圖從資料庫中獲取事件 ID 對應的事件。
    _, err = models.GetEventById(eventId)
    if err != nil {
        context.JSON(500, gin.H{"message": "Could not fetch event"})
        return
    }

    // 創建一個 Event 型別的變數來儲存更新後的事件數據。
    var updatedEvent models.Event
    // 將請求中的 JSON 數據綁定到 updatedEvent 結構體中。
    err = context.ShouldBindJSON(&updatedEvent)
    if err != nil {
        context.JSON(400, gin.H{"message: ": "Invalid request"})
        return
    }

    // 將從 URL 獲取的 eventId 設置到 updatedEvent 結構體的 ID 字段。
    updatedEvent.ID = eventId
    // 調用 Event 結構體的 Update 方法嘗試更新事件到資料庫。
    err = updatedEvent.Update()
    if err != nil {
        context.JSON(500, gin.H{"message ": "Could not update event."})
        return
    }

    context.JSON(200, gin.H{"message": "Event updated!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message: ": "Could not parse event id"})
		return 
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(500, gin.H{"message: ": "Could not fetch event"})
		return 
	}

	err = event.Delete()

	if err != nil {
		context.JSON(500, gin.H{"message: ": "Could not delete event"})
		return
	}

	context.JSON(200, gin.H{"message": "Event delete successfully!"})
}
