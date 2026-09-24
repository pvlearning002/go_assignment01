package service

import (
	"github.com/google/uuid"
	"github.com/pvlearning002/go_assignment01/internal/config"
	"github.com/pvlearning002/go_assignment01/internal/model"
)

//go:generate mockery --name=HealthCheck --filename=health_check_mck.go
type HealthCheck interface {
	GenerateHealthCheck() (*model.HealthCheck, error)
}

type healthCheck struct{}

func NewHealthCheck() HealthCheck {
	return &healthCheck{}
}

func (s *healthCheck) GenerateHealthCheck() (*model.HealthCheck, error) {
	result := &model.HealthCheck{
		Message:     config.OK_STATUS,
		ServiceName: config.BOOKMARK_SERVICE,
		InstanceID:  uuid.New(),
	}
	return result, nil
}
