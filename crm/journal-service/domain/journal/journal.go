package journal

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrInvalidJournalData = errors.New("invalid journal data")
)

type Journal struct {
	id         uuid.UUID
	scheduleID uuid.UUID
	teacherID  uuid.UUID
	date       time.Time
}

func (j *Journal) ID() uuid.UUID {
	return j.id
}

func (j *Journal) ScheduleID() uuid.UUID {
	return j.scheduleID
}

func (j *Journal) Date() time.Time {
	return j.date
}

func (j *Journal) TeacherID() uuid.UUID {
	return j.teacherID
}

// Setters
func (j *Journal) SetDate(date time.Time) error {
	j.date = date
	return j.validate()
}
func (j *Journal) SetScheduleID(id uuid.UUID) {
	j.scheduleID = id
}

func (j *Journal) validate() error {
	if j.date.Equal(time.Time{}) {
		return errors.New("time is empty")
	}

	return nil
}

type UnmarshalJournalArgs struct {
	ID         uuid.UUID
	ScheduleID uuid.UUID
	TeacherID  uuid.UUID
	Date       time.Time
}

func UnmarshalJournal(args UnmarshalJournalArgs) (Journal, error) {
	j := Journal{
		id:         args.ID,
		scheduleID: args.ScheduleID,
		teacherID:  args.TeacherID,
		date:       args.Date,
	}
	if err := j.validate(); err != nil {
		return Journal{}, fmt.Errorf("%w: %s", ErrInvalidJournalData, err.Error())
	}

	return j, nil
}
