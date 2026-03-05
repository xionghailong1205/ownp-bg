package handler

import (
	"net/http"

	"ownp-bg/internal/service"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	healthService service.HealthService
}

func NewHealthHandler(healthService service.HealthService) *HealthHandler {
	return &HealthHandler{healthService: healthService}
}

func (h *HealthHandler) GetStatus(c *gin.Context) {
	status, err := h.healthService.Status(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "database query failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        status.Status,
		"database_time": status.DatabaseTime,
	})
}
