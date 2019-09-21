package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/tobi007/angular-go-serve/config"
	"gopkg.in/go-playground/assert.v1"
)

func TestHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Init("development")

	testRouter := newRouter()

	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
}
