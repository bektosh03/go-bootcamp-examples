package journal

import (
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

type UnmarshalEntryArgs struct {
	JournalID  uuid.UUID
	ScheduleID uuid.UUID
	StudentID  uuid.UUID
	Attended   bool
	Mark       int32
	Date       time.Time
}

func UnmarshalEntry(arg UnmarshalEntryArgs)  {
	return 
}
