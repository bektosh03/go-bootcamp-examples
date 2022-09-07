package service

import (
	"context"
	"schedule-service/domain/schedule"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	CreateSchedule(ctx context.Context, sch schedule.Schedule) error
	GetSchedule(ctx context.Context, scheduleId uuid.UUID) (schedule.Schedule, error)
	UpdateSchedule(ctx context.Context, sch schedule.Schedule) error
	DeleteSchedule(ctx context.Context, scheduleId uuid.UUID) error
	GetFullScheduleForGroup(ctx context.Context, groupId uuid.UUID) ([]schedule.Schedule, error)
	GetScheduleForGroup(ctx context.Context, groupId uuid.UUID, weekday time.Weekday) ([]schedule.Schedule, error)
	GetFullScheduleForTeacher(ctx context.Context, teacherId uuid.UUID) ([]schedule.Schedule, error)
	GetScheduleForTeacher(ctx context.Context, teacherId uuid.UUID, weekday time.Weekday) ([]schedule.Schedule, error)
}
