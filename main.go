package main

import (
	"fmt"

	"github.com/ank809/Event-Management-golang~/authentication"
	"github.com/ank809/Event-Management-golang~/controllers"
	"github.com/ank809/Event-Management-golang~/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	eventroutes := r.Group("/events")
	eventroutes.Use(middlewares.AuthMiddleware())
	eventroutes.Use(middlewares.Authorize("EventManager"))
	{
		eventroutes.POST("/addevent", controllers.AddEvent)
		eventroutes.GET("/deleteevent/:id", controllers.DeleteEvents)
		eventroutes.POST("/updateevent/:id", controllers.UpdateEvent)
	}
	r.POST("/signup", authentication.SignUp)
	r.GET("/loginattendee", authentication.LoginAttendee)
	r.GET("/loginmanager", authentication.EventManagerLogin)
	r.GET("/readevents", controllers.ReadAllEvents)
	r.GET("/registerevent/:id", controllers.RegisterEvent)
	if err := r.Run(":8081"); err != nil {

		fmt.Println(err)
	}

}
