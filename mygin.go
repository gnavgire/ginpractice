package main

import (
	"fmt"
	"gangin/app"
	"github.com/gin-gonic/gin"
)

// This is where you create a gin.Default() and add routes to it
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/cars/:type", app.CarsByType)

	return router
}

func main() {
	fmt.Println("Starting Main")
	router := SetupRouter()
	router.Run(":8080")
}
