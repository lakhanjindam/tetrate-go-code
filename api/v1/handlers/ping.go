package v1handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1models "github.com/tetrate-go-code/api/v1/models"
)

// PingHandler godoc
//
//	@Summary		Ping API v1 server
//	@Description	check api v1 server is up and running
//	@Tags			v1
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	v1models.PingResponse	"OK"
//	@Router			/api/v1/ping [get]
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, v1models.PingResponse{Message: "pong from Gin API v1"})
}
