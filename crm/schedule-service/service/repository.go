package service

import (
	"context"
	"github.com/google/uuid"
	"schedule-service/domain/schedule"
)

type Repository interface {
	CreateSchedule(ctx context.Context, sch schedule.Schedule) error
	GetSchedule(ctx context.Context, scheduleId uuid.UUID) (schedule.Schedule, error)
	UpdateSchedule(ctx context.Context, sch schedule.Schedule) error
	DeleteSchedule(ctx context.Context, scheduleId uuid.UUID) error
	GetFullScheduleForGroup(ctx context.Context, groupId uuid.UUID) ([]schedule.Schedule, error)
	//GetSpecificDateScheduleForGroup(ctx context.Context, groupId uuid.UUID, date time.Weekday) ([]schedule.Schedule, error)
	GetFullScheduleForTeacher(ctx context.Context, teacherId uuid.UUID) ([]schedule.Schedule, error)
	//GetSpecificDateScheduleForTeacher(ctx context.Context, teacherId uuid.UUID, date time.Weekday) ([]schedule.Schedule, error)
}
