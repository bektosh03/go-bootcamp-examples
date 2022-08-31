package service

import (
	"context"
	"schedule-service/domain/schedule"
)

type Repository interface {
	CreateSchedule(ctx context.Context, sch schedule.Schedule) error
}
