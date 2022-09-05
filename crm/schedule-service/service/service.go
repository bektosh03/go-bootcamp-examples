package service

import (
	"context"
	"github.com/google/uuid"
	"schedule-service/domain/schedule"
)

type Service struct {
	repo Repository
}

func (s Service) CreateSchedule(ctx context.Context, sch schedule.Schedule) error {
	return s.repo.CreateSchedule(ctx, sch)
}

func (s Service) GetSchedule(ctx context.Context, id uuid.UUID) (schedule.Schedule, error) {
	return s.repo.GetSchedule(ctx, id)
}

func (s Service) UpdateSchedule(ctx context.Context, sch schedule.Schedule) error {
	return s.repo.UpdateSchedule(ctx, sch)
}

func (s Service) DeleteSchedule(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteSchedule(ctx, id)
}

func (s Service) GetFullScheduleForGroup(ctx context.Context, groupId uuid.UUID) ([]schedule.Schedule, error) {
	return s.repo.GetFullScheduleForGroup(ctx, groupId)
}
