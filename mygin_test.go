package main

import (
	"fmt"
	//"gangin/app"
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGetcarsByType(t *testing.T) {
	gin.SetMode(gin.TestMode)
	//handler := app.CarsByType

	// router_test.go
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", "/cars/1", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	if resp.Code != http.StatusNotFound {
		t.Fatalf("Expecting status 404 not found : got : %v", resp.Code)
	}
}
