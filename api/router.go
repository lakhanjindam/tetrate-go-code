package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "github.com/tetrate-go-code/api/v1"
	"github.com/tetrate-go-code/config"
)

// NewRouter creates a new gin router with the essential middlewares
func NewRouter() *gin.Engine {
	//gin.Default already attaches recovery and logger middleware
	var router = gin.Default()

	// Set up CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: config.Config.AllowOrigins,
		AllowMethods: []string{"GET", "POST"},
		MaxAge:       12 * time.Hour,
	}))

	api := router.Group("/api")
	{

		// v1 APIs
		{
			v1API := api.Group("/v1")
			v1.AddV1Routes(v1API)
		}
	}
	return router
}
