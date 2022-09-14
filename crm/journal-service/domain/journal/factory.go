package journal

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
func (f Factory) NewJournal(scheduleId, teacherID uuid.UUID, date time.Time) (Journal, error) {
	j := Journal{
		id:         f.idGenerator.GenerateUUID(),
		scheduleID: scheduleId,
		teacherID:  teacherID,
		date:       date,
	}
	if err := j.validate(); err != nil {
		return Journal{}, err
	}

	return j, nil
}
