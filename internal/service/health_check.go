package service

import (
	"github.com/google/uuid"
	"github.com/pvlearning002/go_assignment01/internal/config"
	"github.com/pvlearning002/go_assignment01/internal/model"
)

type HealthCheck interface {
	GenerateHealthCheck() *model.HealthCheck
}

type healthCheck struct{}

func NewHealthCheck() HealthCheck {
	return &healthCheck{}
}

func getInstanceID() uuid.UUID {
	id_str := config.HEALTH_CHECK_ID_DEFAULT
	result, err := uuid.Parse(id_str)
	if err != nil {
		result = uuid.New()
	}
	return result
}

func (s *healthCheck) GenerateHealthCheck() *model.HealthCheck {
	return &model.HealthCheck{
		Message:     config.OK_STATUS,
		ServiceName: config.BOOKMARK_SERVICE,
		InstanceID:  getInstanceID(),
	}
}
