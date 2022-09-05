package service

import (
	"context"
	"github.com/google/uuid"
	"schedule-service/domain/schedule"
)

type Repository interface {
	CreateSchedule(ctx context.Context, sch schedule.Schedule) error
	GetSchedule(ctx context.Context, id uuid.UUID) (schedule.Schedule, error)
	UpdateSchedule(ctx context.Context, sch schedule.Schedule) error
	DeleteSchedule(ctx context.Context, id uuid.UUID) error
	GetFullScheduleForGroup(ctx context.Context, groupId uuid.UUID) ([]schedule.Schedule, error)
}
