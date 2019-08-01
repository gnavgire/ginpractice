package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"net/http"
	"ginpractice/app"
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
	router.POST("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "200OK")
	})
	//router.Run(":8080")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown: ", err)
	}

	fmt.Println("Server exiting")
}
