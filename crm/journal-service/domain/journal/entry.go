package journal

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	journalID  uuid.UUID
	scheduleID uuid.UUID
	studentID  uuid.UUID
	attended   bool
	mark       int32
	date       time.Time
}

func (e Entry) JournalID() uuid.UUID {
	return e.journalID
}

func (e Entry) ScheduleID() uuid.UUID {
	return e.scheduleID
}

func (e Entry) StudentID() uuid.UUID {
	return e.studentID
}

func (e Entry) Attended() bool {
	return e.attended
}

func (e Entry) Mark() int32 {
	return e.mark
}

func (e Entry) Date() time.Time {
	return e.date
}

func (e Entry) validate() error {
	if !e.attended && e.mark > 0 {
		return errors.New("unattended student cannot have marks")
	}

	return nil
}

type UnmarshalEntryArgs struct {
	JournalID  uuid.UUID
	ScheduleID uuid.UUID
	Date       time.Time
	StudentID  uuid.UUID
	Attended   bool
	Mark       int32
}

func UnmarshalEntry(arg UnmarshalEntryArgs) (Entry, error) {
	e := Entry{
		journalID:  arg.JournalID,
		scheduleID: arg.ScheduleID,
		studentID:  arg.StudentID,
		attended:   arg.Attended,
		mark:       arg.Mark,
		date:       arg.Date,
	}

	if err := e.validate(); err != nil {
		return Entry{}, err
	}

	return e, nil
}
