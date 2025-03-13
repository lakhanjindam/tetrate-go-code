package v1handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	// create gin server in test mode
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/ping", PingHandler)

	req, err := http.NewRequest("GET", "/ping", nil)

	assert.NoError(t, err)

	// create a response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	want := `{"message":"pong from Gin API v1"}`
	got := w.Body.String()

	assert.Equal(t, want, got)
}
