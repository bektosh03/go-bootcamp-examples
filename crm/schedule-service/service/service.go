package service

import (
	"context"
	"fmt"
	"schedule-service/domain/schedule"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateSchedule(ctx context.Context, sch schedule.Schedule) error {
	return s.repo.CreateSchedule(ctx, sch)
}

func (s Service) GetSchedule(ctx context.Context, scheduleId uuid.UUID) (schedule.Schedule, error) {
	return s.repo.GetSchedule(ctx, scheduleId)
}

func (s Service) UpdateSchedule(ctx context.Context, sch schedule.Schedule) error {
	return s.repo.UpdateSchedule(ctx, sch)
}

func (s Service) DeleteSchedule(ctx context.Context, scheduleId uuid.UUID) error {
	return s.repo.DeleteSchedule(ctx, scheduleId)
}

func (s Service) GetFullScheduleForGroup(ctx context.Context, groupId uuid.UUID) ([]schedule.Schedule, error) {
	return s.repo.GetFullScheduleForGroup(ctx, groupId)
}

func (s Service) GetFullScheduleForTeacher(ctx context.Context, teacherId uuid.UUID) ([]schedule.Schedule, error) {
	return s.repo.GetFullScheduleForTeacher(ctx, teacherId)
}

func (s Service) GetSpecificDateScheduleForTeacher(ctx context.Context, teacherID uuid.UUID, date time.Time) ([]schedule.Schedule, error) {
	return s.repo.GetScheduleForTeacher(ctx, teacherID, date.Weekday())
}

func (s Service) GetSpecificDateScheduleForGroup(ctx context.Context, groupID uuid.UUID, date time.Time) ([]schedule.Schedule, error) {
	fmt.Println("date:", date)
	fmt.Println("weekday:", date.Weekday())
	return s.repo.GetScheduleForGroup(ctx, groupID, date.Weekday())
}
