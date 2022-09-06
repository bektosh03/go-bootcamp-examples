package journal

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrInvalidJournalData = errors.New("invalid journal data")
)

type Journal struct {
	id         uuid.UUID
	scheduleId uuid.UUID
	studentId  uuid.UUID
	attended   bool
	mark       int32
}

func (j Journal) Id() uuid.UUID {
	return j.id
}
func (j Journal) ScheduleId() uuid.UUID {
	return j.scheduleId
}
func (j Journal) StudentId() uuid.UUID {
	return j.studentId
}
func (j Journal) Attended() bool {
	return j.attended
}
func (j Journal) Mark() int32 {
	return j.mark
}
func (j Journal) validate() error {
	if j.mark < 0 || j.mark > 5 {
		return fmt.Errorf("%w invalid mark value", ErrInvalidJournalData)
	}
	return nil
}