package app

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	)

func CarsByType(c *gin.Context) {
    fmt.Println("Go Request in Handler...")
    carType := c.Params.ByName("type")
    fmt.Println(carType)
    if carType != "" {

    }
    c.JSON(http.StatusBadRequest, gin.H{"result": "Bad request"})
    return
}

