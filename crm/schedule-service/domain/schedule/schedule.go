package schedule

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidScheduleData = errors.New("invalid data for Schedule")
)

type Schedule struct {
	id           uuid.UUID
	groupID      uuid.UUID
	subjectID    uuid.UUID
	teacherID    uuid.UUID
	weekday      time.Weekday
	lessonNumber int32
}

func (s Schedule) ID() uuid.UUID {
	return s.id
}
func (s Schedule) GroupID() uuid.UUID {
	return s.groupID
}
func (s Schedule) SubjectID() uuid.UUID {
	return s.subjectID
}
func (s Schedule) TeacherID() uuid.UUID {
	return s.teacherID
}
func (s Schedule) Weekday() time.Weekday {
	return s.weekday
}

func (s Schedule) LessonNumber() int32 {
	return s.lessonNumber
}

func (s Schedule) validate() error {
	if s.weekday == time.Saturday || s.weekday == time.Sunday {
		return fmt.Errorf("%w: a lesson cannot be scheduled for saturday and sunday", ErrInvalidScheduleData)
	}
	if s.lessonNumber < 1 || s.lessonNumber > 6 {
		return fmt.Errorf("%w: a lesson number must be between 1 and 6", ErrInvalidScheduleData)
	}

	return nil
}
