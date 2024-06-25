package main

import (
	"fmt"

	"github.com/ank809/Event-Management-golang~/authentication"
	"github.com/ank809/Event-Management-golang~/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/signup", authentication.SignUp)
	r.GET("/loginattendee", authentication.LoginAttendee)
	r.GET("/loginmanager", authentication.EventManagerLogin)
	r.POST("/addevent", controllers.AddEvent)
	r.GET("/deleteevent/:id", controllers.DeleteEvents)
	r.POST("/updateevent/:id", controllers.UpdateEvent)
	if err := r.Run(":8081"); err != nil {

		fmt.Println(err)
	}

}
