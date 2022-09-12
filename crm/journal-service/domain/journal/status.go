package journal

import (
	"errors"

	"github.com/google/uuid"
)

func NewStatus(journalID, studentID uuid.UUID, attended bool, mark int32) (Stat, error) {
	s := Stat{
		journalID: journalID,
		studentID: studentID,
		attended:  attended,
		mark:      mark,
	}

	if err := s.validate(); err != nil {
		return Stat{}, err
	}

	return s, nil
}

type Stat struct {
	journalID uuid.UUID
	studentID uuid.UUID
	attended  bool
	mark      int32
}

func (s Stat) JournalID() uuid.UUID {
	return s.journalID
}

func (s Stat) StudentID() uuid.UUID {
	return s.studentID
}

func (s Stat) Attended() bool {
	return s.attended
}

func (s Stat) Mark() int32 {
	return s.mark
}

func (s Stat) validate() error {
	if s.mark < 0 || s.mark > 5 {
		return errors.New("mark must be between 0 and 5")
	}

	return nil
}
