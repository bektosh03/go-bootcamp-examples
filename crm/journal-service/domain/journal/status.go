package journal

import (
	"errors"

	"github.com/google/uuid"
)

func NewStatus(journalID, studentID uuid.UUID, attended bool, mark int32) (Status, error) {
	s := Status{
		journalID: journalID,
		studentID: studentID,
		attended:  attended,
		mark:      mark,
	}

	if err := s.validate(); err != nil {
		return Status{}, err
	}

	return s, nil
}

type Status struct {
	journalID uuid.UUID
	studentID uuid.UUID
	attended  bool
	mark      int32
}

func (s Status) JournalID() uuid.UUID {
	return s.journalID
}

func (s Status) StudentID() uuid.UUID {
	return s.studentID
}

func (s Status) Attended() bool {
	return s.attended
}

func (s Status) Mark() int32 {
	return s.mark
}

func (s Status) validate() error {
	if s.mark < 0 || s.mark > 5 {
		return errors.New("mark must be between 0 and 5")
	}

	return nil
}
