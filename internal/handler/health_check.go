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
	modelHealthCheck := h.svc.GenerateHealthCheck()
	c.JSON(http.StatusOK, modelHealthCheck)
}
