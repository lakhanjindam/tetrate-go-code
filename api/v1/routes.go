package v1

import (
	"github.com/gin-gonic/gin"
	v1handlers "github.com/tetrate-go-code/api/v1/handlers"
)

// AddV1Routes adds the v1 routes to the router
func AddV1Routes(v1 *gin.RouterGroup) {
	// Ping route
	v1.GET("/ping", v1handlers.PingHandler)
	v1.GET("/country", v1handlers.GetCountryHandler)
}
