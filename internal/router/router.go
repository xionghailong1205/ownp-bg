package router

import (
	"net/http"

	"ownp-bg/internal/handler"
	"ownp-bg/internal/openapi"

	"github.com/gin-gonic/gin"
)

func New(healthHandler *handler.HealthHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/openapi.json", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(openapi.SpecJSON))
	})

	v1 := r.Group("/api/v1")
	v1.GET("/health", healthHandler.GetStatus)

	return r
}
