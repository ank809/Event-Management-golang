package main

import (
	"fmt"

	"github.com/ank809/Event-Management-golang~/authentication"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/signup", authentication.SignUp)
	r.GET("/loginattendee", authentication.LoginAttendee)
	if err := r.Run(":8081"); err != nil {
		fmt.Println(err)
	}

}
