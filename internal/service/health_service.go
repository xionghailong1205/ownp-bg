package service

import (
	"context"
	"time"

	"ownp-bg/internal/repository"
)

type HealthStatus struct {
	Status       string
	DatabaseTime time.Time
}

type HealthService interface {
	Status(ctx context.Context) (HealthStatus, error)
}

type healthService struct {
	timeRepo repository.TimeRepository
}

func NewHealthService(timeRepo repository.TimeRepository) HealthService {
	return &healthService{timeRepo: timeRepo}
}

func (s *healthService) Status(ctx context.Context) (HealthStatus, error) {
	dbTime, err := s.timeRepo.CurrentTime(ctx)
	if err != nil {
		return HealthStatus{}, err
	}

	return HealthStatus{
		Status:       "ok",
		DatabaseTime: dbTime,
	}, nil
}
