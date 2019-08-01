package app

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	)

func CarsByType(c *gin.Context) {
    fmt.Println("Go Request in Handler...")
    carType := c.Params.ByName("type")
    //fmt.Println(carType)
    bhp := c.Query("bhp")
    fmt.Println(bhp)
    if carType != "" {

	    c.JSON(http.StatusOK, gin.H{"BMW": "850BHP ", "Audi" : "800BHP", "Merc" : "1000BHP", "Jaguar" : "1100BHP"})
    } else {
    	c.JSON(http.StatusBadRequest, gin.H{"result": "Success "})
   }
    return
}

