package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pvlearning002/go_assignment01/internal/service"
)

type HealthCheck interface {
	GenerateHealthCheck(c *gin.Context)
}

type healthCheck struct {
	svc service.HealthCheck
}

func NewHealthCheck(svc service.HealthCheck) HealthCheck {
	return &healthCheck{svc: svc}
}

func (h *healthCheck) GenerateHealthCheck(c *gin.Context) {
	modelHealthCheck, err := h.svc.GenerateHealthCheck()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, modelHealthCheck)
}
