package schedule

import (
	"time"

	"github.com/bektosh03/crmcommon/id"
	"github.com/google/uuid"
)

type Factory struct {
	idGenerator id.IGenerator
}

func NewFactory(idGenerator id.IGenerator) Factory {
	return Factory{
		idGenerator: idGenerator,
	}
}

func (f Factory) NewSchedule(
	groupID, subjectID, teacherID uuid.UUID, weekday time.Weekday, lessonNumber int32,
) (Schedule, error) {
	s := Schedule{
		id:           f.idGenerator.GenerateUUID(),
		groupID:      groupID,
		subjectID:    subjectID,
		teacherID:    teacherID,
		weekday:      weekday,
		lessonNumber: lessonNumber,
	}

	if err := s.validate(); err != nil {
		return Schedule{}, err
	}

	return s, nil
}
