package main

import (
	"github.com/gin-gonic/gin"
    "GoWorkSpace/goCLinic/handler"
)

func main() {
	router := gin.Default()

	router.POST("/clinics", handler.CreateClinicHandler)
	//router.POST("/tickets", createTicketHandler)
	//router.GET("/tickets/:id", showTicketHandler)

	router.Run(":8080")
}
