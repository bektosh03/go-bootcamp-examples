package service

import (
	"context"
	"schedule-service/domain/schedule"
)

type Service struct {
	repo Repository
}

func (s Service) CreateSchedule(ctx context.Context, sch schedule.Schedule) error {
	return s.repo.CreateSchedule(ctx, sch)
}
