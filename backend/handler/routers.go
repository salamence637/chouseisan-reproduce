package handler

import "github.com/gin-gonic/gin"

func SetupEventRoutes(router *gin.Engine, event_handler *EventHandler) {
	/*
		router.GET("/chouseisan/events", GetEventByHash)
		router.POST("chouseisan/events", CreateEvent)
	*/
	router.POST("/event", event_handler.CreateEventHandler)
	router.DELETE("/event/:eventID", event_handler.DeleteEventHandler)
	router.GET("/event/basic/:eventID", event_handler.GetEventBasicHandler)
	router.PUT("/event/editTitleDetail/:eventID", event_handler.EditTitleDetailHandler)
	router.GET("/event/timeslots/:eventID", event_handler.GetTimeslotsHandler)
	router.PUT("event/deleteTimeslots/:eventID", event_handler.DeleteTimeslotsHandler)
	router.PUT("event/addTimeslots/:eventID", event_handler.AddTimeslotsHandler)
	router.GET("event/exist/:eventID", event_handler.CheckEventExistsHandler)
	router.GET("event/isCreatedBySelf/:eventID", event_handler.IsCreatedBySelfHandler)
	router.POST("attendance/:eventID", event_handler.AddAttendanceHandler)
	router.GET("attendance/:eventID", event_handler.GetAttendanceHandler)
	router.PUT("attendance/:eventID", event_handler.ModifyAttendanceHandler)
	router.PUT("/event/editDue/:eventID", event_handler.EditDueHandler)
}
