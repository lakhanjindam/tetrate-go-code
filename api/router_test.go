package api

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/tetrate-go-code/config"
)

func TestNewRouter(t *testing.T) {
	err := config.LoadConfig("../config")
	assert.NoError(t, err)

	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	router := NewRouter()

	assert.NotNil(t, router)

	assert.NotNil(t, router.Group("/v1"))

	want := []string{
		"GET /api/v1/ping",
		"GET /api/v1/country",
	}

	for _, route := range router.Routes() {
		route := fmt.Sprintf("%s %s", route.Method, route.Path)
		assert.Contains(t, want, route)
	}
}
