package journal

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrInvalidJournalData = errors.New("invalid journal data")
)

type Journal struct {
	id         uuid.UUID
	scheduleId uuid.UUID
	date       time.Time
}

func (j Journal) ID() uuid.UUID {
	return j.id
}

func (j Journal) ScheduleID() uuid.UUID {
	return j.scheduleId
}

func (j Journal) Date() time.Time {
	return j.date
}

// Setters
func (j *Journal) SetDate(date time.Time) error {
	j.date = date
	return j.validate()
}
func (j *Journal) SetScheduleID(id uuid.UUID) {
	j.scheduleId = id
}

func (j Journal) validate() error {
	if j.date.Equal(time.Time{}) {
		return errors.New("time is empty")
	}

	return nil
}


type UnmarshalJournalArgs struct {
	ID         uuid.UUID
	ScheduleID uuid.UUID
	Date       time.Time
}

func UnmarshalJournal(args UnmarshalJournalArgs) (Journal, error) {
	j := Journal{
		id:         args.ID,
		scheduleId: args.ScheduleID,
		date:       args.Date,
	}
	if err := j.validate(); err != nil {
		return Journal{}, nil
	}

	return j, nil
}
