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

func (s Service) GetFullScheduleForGroup(ctx context.Context, groupID uuid.UUID) ([]schedule.Schedule, error) {
	schs, err := s.repo.GetFullScheduleForGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return s.sortScheduleForOneDate(schs), nil
}

func (s Service) GetFullScheduleForTeacher(ctx context.Context, teacherID uuid.UUID) ([]schedule.Schedule, error) {
	schs, err := s.repo.GetFullScheduleForTeacher(ctx, teacherID)
	if err != nil {
		return nil, err
	}

	return s.sortScheduleForOneDate(schs), nil
}

func (s Service) GetSpecificDateScheduleForTeacher(ctx context.Context, teacherID uuid.UUID, date time.Time) ([]schedule.Schedule, error) {
	fmt.Println("date!:", date)
	fmt.Println("week!:", date.Weekday())

	schs, err := s.repo.GetScheduleForTeacher(ctx, teacherID, date.Weekday())
	if err != nil {
		return nil, err
	}

	return s.sortScheduleForOneDate(schs), nil
}

func (s Service) GetSpecificDateScheduleForGroup(ctx context.Context, groupID uuid.UUID, date time.Time) ([]schedule.Schedule, error) {
	fmt.Println("date:", date)
	fmt.Println("weekday:", date.Weekday())

	schs, err := s.repo.GetScheduleForGroup(ctx, groupID, date.Weekday())
	if err != nil {
		return nil, err
	}

	return s.sortScheduleForOneDate(schs), nil
}

func (s Service) sortScheduleForOneDate(schs []schedule.Schedule) []schedule.Schedule {
	for i := 0; i < len(schs); i++ {
		for j := i + 1; j < len(schs); j++ {
			if schs[i].Weekday() > schs[j].Weekday() {
				schs[i], schs[j] = schs[j], schs[i]
			} else if schs[i].Weekday() == schs[j].Weekday() {
				if schs[i].LessonNumber() > schs[j].LessonNumber() {
					schs[i], schs[j] = schs[j], schs[i]
				}
			}
		}
	}

	return schs
}
