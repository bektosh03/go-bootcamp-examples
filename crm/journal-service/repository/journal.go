package repository

import "github.com/google/uuid"

type Journal struct {
	ID         uuid.UUID `json:"id"`
	ScheduleID uuid.UUID `json:"schedule_id"`
	StudentID  uuid.UUID `json:"student_id"`
	Attended   bool      `json:"attended"`
	Mark       int32     `json:"mark"`
}
