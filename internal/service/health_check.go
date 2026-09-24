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

func (s *healthCheck) GetInstanceID() uuid.UUID {
	return uuid.New()
}

func (s *healthCheck) GenerateHealthCheck() (*model.HealthCheck, error) {
	cfg := config.GetConfig()
	result := &model.HealthCheck{
		Message:     cfg.OKStatus,
		ServiceName: cfg.BookmarkService,
		InstanceID:  s.GetInstanceID(),
	}
	return result, nil
}
